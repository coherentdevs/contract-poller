package evm_client

import (
	"context"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	protos "github.com/coherent-api/contract-poller/protos/go/protos/evm/node_client"
	"github.com/coherent-api/contract-poller/protos/go/protos/shared"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/http"
	"time"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/coherent-api/contract-poller/shared/service_framework"
)

type nodeClient struct {
	url          string
	parsedClient *ethclient.Client

	httpClient *http.Client
	config     *Config
	blockchain constants.Blockchain
}

func getNode(config *Config, blockchain constants.Blockchain) string {
	switch blockchain {
	case constants.Ethereum:
		return config.EthNodeRPC
	case constants.Optimism:
		return config.OptNodeRPC
	case constants.Polygon:
		return config.PolyNodeRPC
	case constants.Goerli:
		return config.GoerliNodeRPC
	}
	return ""
}

func NewClient(config *Config, manager *service_framework.Manager) (*nodeClient, error) {
	url := getNode(config, config.Blockchain)
	parsedClient, err := ethclient.Dial(url)
	if err != nil {
		manager.Logger().Fatal(err)
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: time.Second * 300,
	}

	return &nodeClient{
		url:          url,
		httpClient:   httpClient,
		parsedClient: parsedClient,
		config:       config,
		blockchain:   config.Blockchain,
	}, nil
}

func MustNewClient(config *Config, manager *service_framework.Manager) *nodeClient {
	client, err := NewClient(config, manager)
	if err != nil {
		manager.Logger().With(err).Fatal("Failed to instantiate node client")
	}
	return client
}

func (c *nodeClient) GetContract(address string) (*models.Contract, error) {
	ctx := context.Background()
	contractReq := &protos.GetContractRequest{
		Address:    address,
		Blockchain: shared.Blockchain(shared.Blockchain_value[c.blockchain.GetSymbol()]),
	}

	contractResp, err := c.GetContractFromNode(ctx, contractReq)
	if err != nil {
		return nil, err
	}
	contract := &models.Contract{
		Address:    address,
		Blockchain: c.blockchain,
		Name:       contractResp.Name,
		Symbol:     contractResp.Symbol,
		//Standard:   contractResp.Standard,
		Decimals: contractResp.Decimals,
	}
	return contract, nil
}
