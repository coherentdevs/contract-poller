package abi_client

import (
	"github.com/caarlos0/env/v7"
	"github.com/coherentopensource/go-service-framework/util"
	"time"

	"github.com/datadaodevs/go-service-framework/constants"
	"github.com/nanmu42/etherscan-api"
)

type Config struct {
	EtherscanAPIKey           string               `env:"ETHERSCAN_API_KEY,required"`
	AbiClientRateMilliseconds time.Duration        `env:"ABI_CLIENT_RATE" envDefault:"100ms"`
	AbiClientRateRequests     int                  `env:"ABI_CLIENT_RATE_REQUESTS" envDefault:"200"`
	EtherscanErrorSleep       time.Duration        `env:"ETHERSCAN_ERROR_SLEEP" envDefault:"1000ms"`
	EtherscanNetwork          etherscan.Network    `env:"ETHERSCAN_NETWORK" envDefault:"api"`
	Blockchain                constants.Blockchain `env:"BLOCKCHAIN,required"`
	HTTPRetries               int                  `env:"HTTP_RETRIES" envDefault:"3"`
}

func MustParseConfig(logger util.Logger) *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		logger.Fatalf("could not parse Base driver config: %v", err)
	}

	return &cfg
}
