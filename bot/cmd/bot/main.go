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

	repo, err := psql.New(conf.PostgresDSN)
	if err != nil {
		sugar.Fatal(err)
	}

	ton, err := setupTON(ctx, conf.TON)
	if err != nil {
		err = fmt.Errorf("setup ton: %w", err)
		sugar.Fatal(err)
	}

	ghClient, err := setupGitHubClient(conf.GitHub)
	if err != nil {
		err = fmt.Errorf("setup github client: %w", err)
		sugar.Fatal(err)
	}

	svc := service.New(&service.Deps{
		TON:          ton,
		GitHubClient: ghClient,
		Repository:   repo,
	})

	handler, err := handlers.New(svc, conf.GitHub.WebhookSecret, sugar)
	if err != nil {
		sugar.Fatal(err)
	}
	srv := server.New(handler.HTTP(), conf.ServerPort)

	log.Println("Listening...")
	if err := srv.Run(ctx); err != nil {
		err = fmt.Errorf("run: %w", err)
		sugar.Fatal(err)
	}
}

func setupTON(ctx context.Context, config configs.TON) (*ton.TON, error) {
	pool := liteclient.NewConnectionPool()
	err := pool.AddConnectionsFromConfigUrl(ctx, config.ConfigURL)
	if err != nil {
		return nil, fmt.Errorf("add ton connection: %w", err)
	}
	tonClient := tonUtils.NewAPIClient(pool)

	seed := strings.Split(config.WalletSeed, " ")
	tonWallet, err := wallet.FromSeed(tonClient, seed, wallet.V3)
	if err != nil {
		return nil, fmt.Errorf("wallet from seed: %w", err)
	}

	return ton.New(tonClient, tonWallet, config.RouterContract), nil
}

func setupGitHubClient(config configs.GitHub) (*http.Client, error) {
	itr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, config.AppID, config.PkPath)
	if err != nil {
		return nil, fmt.Errorf("create transport: %w", err)
	}

	return &http.Client{Transport: itr}, nil
}
