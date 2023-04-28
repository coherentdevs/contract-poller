package ethereum

import (
	"github.com/caarlos0/env/v7"
	"github.com/coherentopensource/go-service-framework/util"
	"github.com/datadaodevs/go-service-framework/constants"
)

// Config stores configurable properties of the driver
type Config struct {
	MaxRetries int                  `env:"HTTP_MAX_RETRIES" envDefault:"5"`
	Blockchain constants.Blockchain `env:"BLOCKCHAIN,required"`
}

// MustParseConfig uses env.Parse to initialize config with environment variables
func MustParseConfig(logger util.Logger) *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		logger.Fatalf("could not parse Ethereum driver config: %v", err)
	}

	return &cfg
}
