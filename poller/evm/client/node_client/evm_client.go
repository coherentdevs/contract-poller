package node_client

import (
	"context"
	"net/http"
	"time"

	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/coherent-api/contract-poller/shared/service_framework"
)

type evmClient struct {
	url          string
	parsedClient *ethclient.Client

	httpClient *http.Client
	config     *Config
	blockchain constants.Blockchain
}

type GetContractRequest struct {
	Address    string
	Blockchain constants.Blockchain
}

type GetContractResponse struct {
	Address  string
	Name     string
	Symbol   string
	Type     AddressType
	Decimals int32
	Bytecode []byte
}

type AddressType int32

const (
	AddressType_USER     AddressType = 0
	AddressType_CONTRACT AddressType = 1
	AddressType_ERC20    AddressType = 2
	AddressType_ERC721   AddressType = 3
	AddressType_ERC1155  AddressType = 4
)

func getNode(config *Config, blockchain constants.Blockchain) string {
	switch blockchain {
	case constants.Ethereum:
		return config.EthNodeRPC
	case constants.Optimism:
		return config.OptNodeRPC
	case constants.Polygon:
		return config.PolyNodeRPC
	}
	return ""
}

func NewClient(config *Config, manager *service_framework.Manager) (*evmClient, error) {
	url := getNode(config, config.Blockchain)
	parsedClient, err := ethclient.Dial(url)
	if err != nil {
		manager.Logger().Fatal(err)
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: time.Second * 300,
	}

	return &evmClient{
		url:          url,
		httpClient:   httpClient,
		parsedClient: parsedClient,
		config:       config,
		blockchain:   config.Blockchain,
	}, nil
}

func MustNewClient(config *Config, manager *service_framework.Manager) *evmClient {
	client, err := NewClient(config, manager)
	if err != nil {
		manager.Logger().With(err).Fatal("Failed to instantiate node client")
	}
	return client
}

func (c *evmClient) GetContract(address string) (*models.Contract, error) {
	ctx := context.Background()
	contractReq := &GetContractRequest{
		Address:    address,
		Blockchain: c.blockchain,
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
