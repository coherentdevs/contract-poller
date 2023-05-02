package abi_client

import (
	"context"
	"fmt"
	"github.com/coherentopensource/go-service-framework/rate_limiter"
	"github.com/nanmu42/etherscan-api"
)

type ethereumClient struct {
	Client      *etherscan.Client
	RateLimiter *rate_limiter.RateLimitedClient
}

func NewEthereumABIClient(cfg *Config) (*ethereumClient, error) {

	if cfg.EtherscanNetwork == "" {
		return nil, fmt.Errorf("etherscan network is not defined")
	}

	if cfg.EtherscanAPIKey == "" {
		return nil, fmt.Errorf("etherscan api key is not defined")
	}
	rateLimiter := rate_limiter.NewClient(cfg.ErrorSleep, cfg.RateIntervalMs, cfg.MaxRateRequests)

	client := etherscan.New(cfg.EtherscanNetwork, cfg.EtherscanAPIKey)
	return &ethereumClient{
		Client:      client,
		RateLimiter: rateLimiter,
	}, nil
}

// ContractSource returns the contract source code for a given contract address
func (e *ethereumClient) ContractSource(ctx context.Context, contractAddress string) (*etherscan.ContractSource, error) {
	var contractSources []etherscan.ContractSource
	var err error
	rateErr := e.RateLimiter.Exec(ctx, func() error {
		contractSources, err = e.Client.ContractSource(contractAddress)
		if err != nil {
			return err
		}
		return nil
	})
	if rateErr != nil {
		return nil, rateErr
	}
	return &contractSources[0], nil
}
