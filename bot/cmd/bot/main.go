package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/deed-labs/gittips/bot/configs"
	"github.com/deed-labs/gittips/bot/internal/github"
	"github.com/deed-labs/gittips/bot/internal/handlers"
	"github.com/deed-labs/gittips/bot/internal/repository/psql"
	"github.com/deed-labs/gittips/bot/internal/server"
	"github.com/deed-labs/gittips/bot/internal/service"
	"github.com/joho/godotenv"
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

	svc := service.New(&service.Deps{
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
