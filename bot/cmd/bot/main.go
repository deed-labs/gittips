package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/deed-labs/gittips/bot/configs"
	"github.com/deed-labs/gittips/bot/internal/github"
	"github.com/deed-labs/gittips/bot/internal/handlers"
	"github.com/deed-labs/gittips/bot/internal/repository/psql"
	"github.com/deed-labs/gittips/bot/internal/server"
	"github.com/deed-labs/gittips/bot/internal/service"
	"github.com/deed-labs/gittips/bot/internal/ton"
	"github.com/joho/godotenv"
	"github.com/xssnick/tonutils-go/liteclient"
	tonUtils "github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := godotenv.Load(); err != nil {
		sugar.Warn("no .env file found")
	}

	conf, err := configs.LoadConfig()
	if err != nil {
		sugar.Fatal(err)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-exit
		cancel()
	}()

	// Connect to TON

	pool := liteclient.NewConnectionPool()
	err = pool.AddConnection(ctx, conf.TON.URL, conf.TON.ServerKey)
	if err != nil {
		err = fmt.Errorf("add ton connection: %w", err)
		sugar.Fatal(err)
	}
	tonClient := tonUtils.NewAPIClient(pool)
	seed := strings.Split(conf.TON.WalletSeed, " ")
	tonWallet, err := wallet.FromSeed(tonClient, seed, wallet.V3)
	if err != nil {
		err = fmt.Errorf("wallet from seed: %w", err)
		sugar.Fatal(err)
	}

	ton := ton.New(tonClient, tonWallet, conf.TON.RouterContract)

	repo, err := psql.New(conf.PostgresDSN)
	if err != nil {
		sugar.Fatal(err)
	}

	svc := service.New(&service.Deps{
		TON:        ton,
		Repository: repo,
	})

	itr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, conf.GitHubAppID, conf.GitHubAppPkPath)
	if err != nil {
		sugar.Fatal(err)
	}

	gh := github.New(conf.GitHubWebhookSecret, &http.Client{Transport: itr}, svc)
	ghHandler, err := github.NewHandler(gh, sugar)
	if err != nil {
		sugar.Fatal(err)
	}

	handler := handlers.New(ghHandler.Handle, svc)
	srv := server.New(handler.HTTP(), conf.ServerPort)

	log.Println("Listening...")
	if err := srv.Run(ctx); err != nil {
		err = fmt.Errorf("run: %w", err)
		sugar.Fatal(err)
	}
}
