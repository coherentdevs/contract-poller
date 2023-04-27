package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/coherentopensource/chain-interactor/client/node"
	"github.com/coherentopensource/contract-poller/poller/pkg/cache"
	"github.com/coherentopensource/go-service-framework/contract_poller"
	"github.com/coherentopensource/go-service-framework/manager"
	"github.com/coherentopensource/go-service-framework/util"
	"time"
)

type Environment string

type Config struct {
	Env manager.Environment `env:"ENV" envDefault:"local"`

	FetcherPoolThrottleBandwidth int           `env:"FETCHER_POOL_THROTTLER_BANDWIDTH" envDefault:"100"`
	FetcherPoolThrottleDuration  time.Duration `env:"FETCHER_POOL_THROTTLER_DURATION" envDefault:"60s"`

	FetcherPoolBandwidth     int `env:"FETCHER_POOL_BANDWIDTH" envDefault:"100"`
	AccumulatorPoolBandwidth int `env:"ACCUMULATOR_POOL_BANDWIDTH" envDefault:"100"`
	WriterPoolBandwidth      int `env:"WRITER_POOL_BANDWIDTH" envDefault:"100"`

	HttpServePort  string `env:"HTTP_SERVE_PORT" envDefault:":8000"`
	PprofServePort string `env:"PPROF_SERVE_PORT" envDefault:":6060"`

	Cache  cache.Config
	Poller contract_poller.Config
	Node   node.Config
}

func MustNewConfig(logger util.Logger) *Config {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		logger.Fatalf("Failed to parse service config: %v", err)
	}

	return &cfg
}
