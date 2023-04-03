package abi_client

import (
	"context"
	"fmt"

	"time"

	"github.com/pkg/errors"

	"github.com/datadaodevs/go-service-framework/constants"
	"github.com/datadaodevs/go-service-framework/retry"
	"github.com/nanmu42/etherscan-api"
)

var (
	ErrEtherscanServerNotOK = errors.New("etherscan server: NOTOK")
)

type ethereumClient struct {
	HTTPRetries int

	Client *etherscan.Client
}

func NewEthereum(cfg *Config) (*ethereumClient, error) {

	if cfg.EtherscanNetwork == "" {
		return nil, fmt.Errorf("etherscan network is not defined")
	}

	if cfg.EtherscanAPIKey == "" {
		return nil, fmt.Errorf("etherscan api key is not defined")
	}

	client := etherscan.New(cfg.EtherscanNetwork, cfg.EtherscanAPIKey)
	return &ethereumClient{
		Client:      client,
		HTTPRetries: cfg.HTTPRetries,
	}, nil
}

/**
 * Get the contract source code from etherscan
 */
func (e *ethereumClient) ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error) {
	var contractSources []etherscan.ContractSource

	err := retry.Exec(e.HTTPRetries, func(attempt int) (bool, error) {
		var err error
		contractSources, err = e.Client.ContractSource(contractAddress)

		isRetriableErr := (err != nil && errors.Is(err, ErrEtherscanServerNotOK))

		if isRetriableErr {
			exp := time.Duration(2 ^ (attempt - 1))
			time.Sleep(exp * time.Second)
		}
		return (attempt < e.HTTPRetries && isRetriableErr), err
	})

	if err != nil {
		return etherscan.ContractSource{}, fmt.Errorf("failed to get contract resource for contract: %s: %v", contractAddress, err)
	}

	return contractSources[0], nil
}
