package abi_client

import (
	"context"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/coherent-api/contract-poller/shared/service_framework"
	"github.com/nanmu42/etherscan-api"
)

type Client interface {
	ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error)
}

func MustNew(cfg *Config, manager *service_framework.Manager) Client {
	var client Client
	var err error

	switch cfg.Blockchain {
	case constants.Ethereum:
		client, err = NewEthereum(cfg)
		if err != nil {
			manager.Logger().With(err).Fatal("Failed to instantiate eth abi client")
		}
	case constants.Polygon:
		client, err = NewPolygon(cfg)
		if err != nil {
			manager.Logger().With(err).Fatal("Failed to instantiate polygon abi client")
		}
	}
	return client
}
