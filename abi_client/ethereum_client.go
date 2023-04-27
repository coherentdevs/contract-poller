package abi_client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/coherentopensource/go-service-framework/retry"
	"github.com/nanmu42/etherscan-api"
)

var (
	ErrEtherscanServerNotOK = errors.New("etherscan server: NOTOK")
)

type ethereumClient struct {
	HTTPRetries int
	Client      *etherscan.Client
}

func NewEthereumABIClient(cfg *Config) (*ethereumClient, error) {

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

// ContractSource returns the contract source code for a given contract address
func (e *ethereumClient) ContractSource(ctx context.Context, contractAddress string) (*etherscan.ContractSource, error) {
	var contractSources []etherscan.ContractSource
	var err error
	if retryErr := retry.Exec(e.HTTPRetries, func() error {
		contractSources, err = e.Client.ContractSource(contractAddress)
		if errors.Is(err, ErrEtherscanServerNotOK) {
			return err
		}
		return nil
	}, nil); retryErr != nil {
		return nil, retryErr
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get contract resource for contract: %s: %v", contractAddress, err)
	}

	return &contractSources[0], nil
}
