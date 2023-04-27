package abi_client

import (
	"context"
	"github.com/nanmu42/etherscan-api"
)

type Client interface {
	ContractSource(ctx context.Context, contractAddress string) (*etherscan.ContractSource, error)
}
