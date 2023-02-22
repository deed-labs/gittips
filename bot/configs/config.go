package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerPort          int    `env:"SERVER_PORT"`
	PostgresDSN         string `env:"POSTGRES_DSN"`
	GitHubAppID         int64  `env:"GITHUB_APP_ID"`
	GitHubAppPkPath     string `env:"GITHUB_APP_PK_PATH"`
	GitHubWebhookSecret string `env:"GITHUB_WEBHOOK_SECRET"`
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return c, nil
}
