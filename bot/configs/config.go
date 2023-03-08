package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerPort  int    `env:"SERVER_PORT"`
	PostgresDSN string `env:"POSTGRES_DSN"`

	GitHub GitHub
	TON    TON
}

type GitHub struct {
	AppID         int64  `env:"GITHUB_APP_ID"`
	PkPath        string `env:"GITHUB_APP_PK_PATH"`
	WebhookSecret string `env:"GITHUB_WEBHOOK_SECRET"`
}

type TON struct {
	URL            string `env:"TON_URL"`
	RouterContract string `env:"TON_ROUTER_CONTRACT"`
	ServerKey      string `env:"TON_SERVER_KEY"`
	WalletSeed     string `env:"TON_WALLET_SEED"`
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return c, nil
}
