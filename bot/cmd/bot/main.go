package main

import (
	"context"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/deed-labs/openroll/bot/configs"
	"github.com/deed-labs/openroll/bot/internal/github"
	"github.com/deed-labs/openroll/bot/internal/handlers"
	"github.com/deed-labs/openroll/bot/internal/repository/psql"
	"github.com/deed-labs/openroll/bot/internal/server"
	"github.com/deed-labs/openroll/bot/internal/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(fmt.Errorf("no .env file found"))
	}

	conf, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	svc := service.New(&service.Deps{
		Repository: repo,
	})

	itr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, conf.GitHubAppID, conf.GitHubAppPkPath)
	if err != nil {
		log.Fatal(err)
	}

	gh := github.New(conf.GitHubWebhookSecret, &http.Client{Transport: itr}, svc)
	ghHandler, err := gh.Handler()
	if err != nil {
		log.Fatal(err)
	}

	hndl := handlers.New(ghHandler.Handle, svc)
	srv := server.New(hndl.HTTP(), conf.ServerPort)

	log.Println("Listening...")
	if err := srv.Run(ctx); err != nil {
		err = fmt.Errorf("run bridge: %w", err)
		log.Fatal(err)
	}
}
