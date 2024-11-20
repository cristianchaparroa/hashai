package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TheGraphApiKey    string `env:"THE_GRAPH_API_KEY"`
	PolygonRpcURl     string `env:"RPC_URL"`
	PolygonPrivateKey string `env:"PRIVATE_KEY"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return cfg, nil
}
