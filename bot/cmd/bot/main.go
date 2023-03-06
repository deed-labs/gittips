package main

import (
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/deed-labs/gittips/bot/configs"
	"github.com/deed-labs/gittips/bot/internal/handlers"
	"github.com/deed-labs/gittips/bot/internal/repository/psql"
	"github.com/deed-labs/gittips/bot/internal/server"
	"github.com/deed-labs/gittips/bot/internal/service"
	"github.com/deed-labs/gittips/bot/internal/ton"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	ghClient, err := setupGithubClient(conf.Github)
	if err != nil {
		err = fmt.Errorf("setup github client: %w", err)
		sugar.Fatal(err)
	}

	svc := service.New(&service.Deps{
		TON:          ton,
		GithubClient: ghClient,
		Repository:   repo,
	})

	handler, err := handlers.New(svc, conf.Github.WebhookSecret, sugar)
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
	// TODO: uncomment after contract deployed
	//pool := liteclient.NewConnectionPool()
	//err := pool.AddConnection(ctx, config.URL, config.ServerKey)
	//if err != nil {
	//	return nil, fmt.Errorf("add ton connection: %w", err)
	//}
	//tonClient := tonUtils.NewAPIClient(pool)
	//
	//seed := strings.Split(config.WalletSeed, " ")
	//tonWallet, err := wallet.FromSeed(tonClient, seed, wallet.V3)
	//if err != nil {
	//	return nil, fmt.Errorf("wallet from seed: %w", err)
	//}
	//
	//return ton.New(tonClient, tonWallet, config.RouterContract), nil

	return &ton.TON{}, nil
}

func setupGithubClient(config configs.Github) (*http.Client, error) {
	itr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, config.AppID, config.PkPath)
	if err != nil {
		return nil, fmt.Errorf("create transport: %w", err)
	}

	return &http.Client{Transport: itr}, nil
}
