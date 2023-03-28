package abi_client

import (
	"context"
	"fmt"

	"time"

	"github.com/pkg/errors"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/nanmu42/etherscan-api"
)

var (
	ErrEtherscanServerNotOK = errors.New("etherscan server: NOTOK")
)

type EthereumClient struct {
	cfg *Config

	Client     *etherscan.Client
	ErrorSleep time.Duration
}

func NewEthereum(cfg *Config) (*EthereumClient, error) {

	if cfg.EtherscanNetwork == "" {
		return nil, fmt.Errorf("etherscan network is not defined")
	}

	if cfg.EtherscanAPIKey == "" {
		return nil, fmt.Errorf("etherscan api key is not defined")
	}

	client := etherscan.New(cfg.EtherscanNetwork, cfg.EtherscanAPIKey)
	return &EthereumClient{
		Client:     client,
		cfg:        cfg,
		ErrorSleep: cfg.EtherscanErrorSleep,
	}, nil
}

/**
 * Get the contract source code from etherscan
 */
func (r *EthereumClient) ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error) {
	contractSources, err := r.rateLimitedContractSource(ctx, contractAddress, 0, blockchain)
	if err != nil {
		return etherscan.ContractSource{}, err
	}
	return contractSources[0], nil
}

func (r *EthereumClient) rateLimitedContractSource(ctx context.Context, contractAddress string, attemptCount int, blockchain constants.Blockchain) ([]etherscan.ContractSource, error) {
	var contractSources []etherscan.ContractSource
	var err error

	contractSources, err = r.Client.ContractSource(contractAddress)
	if err != nil {
		if errors.Is(err, ErrEtherscanServerNotOK) && attemptCount < 5 {
			time.Sleep(r.ErrorSleep * time.Millisecond)
			return r.rateLimitedContractSource(ctx, contractAddress, attemptCount+1, blockchain)
		}
		return nil, err
	}

	return contractSources, nil
}
