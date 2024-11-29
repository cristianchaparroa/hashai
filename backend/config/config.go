package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TheGraph struct {
		ApiKey string `env:"THE_GRAPH_API_KEY"`
	}

	Polygon struct {
		RpcURl     string `env:"POLYGON_RPC_URL" yaml:"rpc_url"`
		ChainID    int64  `env:"POLYGON_CHAIN_ID" yaml:"chain_id"`
		PrivateKey string `env:"POLYGON_PRIVATE_KEY"`
	} `yaml:"polygon"`

	ReportContract struct {
		Address string `env:"REPORT_ADDRESS" yaml:"address"`
		AbiFile string `env:"REPORT_ABI_FILE" yaml:"abi_file"`
	} `yaml:"report_contract"`

	BlackListContract struct {
		Address string `env:"BLACKLIST_ADDRESS" yaml:"address"`
		AbiFile string `env:"BLACKLIST_ABI_FILE" yaml:"abi_file"`
	} `yaml:"blacklist_contract"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return cfg, nil
}
