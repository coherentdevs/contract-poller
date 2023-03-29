package abi_client

import (
	"context"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/nanmu42/etherscan-api"
)

type Client interface {
	ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error)
}
