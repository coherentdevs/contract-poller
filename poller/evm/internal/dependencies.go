package contract_poller

import (
	"context"

	"github.com/coherent-api/contract-poller/poller/evm/client/abi_client"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/datadaodevs/go-service-framework/util"
	"github.com/nanmu42/etherscan-api"
)

type ABIClient interface {
	ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error)
}

type Database interface {
	GetContractsToBackfill() ([]models.Contract, error)
	UpdateContractsToBackfill(contracts []models.Contract) error
	GetContract(contractAddress string, blockchain constants.Blockchain) (*models.Contract, error)
	GetEventFragmentById(eventId string) (*models.EventFragment, error)
	GetMethodFragmentByID(methodId string) (*models.MethodFragment, error)
	DeleteContractByAddress(address string) error
	DeleteEventFragment(eventFragment *models.EventFragment) error
	DeleteMethodFragment(methodFragment *models.MethodFragment) error
	UpdateContractByAddress(contract *models.Contract) error
	UpdateMethodFragment(methodFragment *models.MethodFragment) error
	UpdateEventFragment(eventFragment *models.EventFragment) error
	UpsertContracts(contracts []models.Contract) (int64, error)
	UpsertEventFragment(eventFragment *models.EventFragment) (int64, error)
	UpsertMethodFragment(methodFragment *models.MethodFragment) (int64, error)
}

type EVMClient interface {
	GetContract(contractAddress string) (*models.Contract, error)
}

func MustNewABIClient(blockchain constants.Blockchain, cfg *abi_client.Config, logger util.Logger) ABIClient {
	var client ABIClient
	var err error

	switch blockchain {
	case constants.Ethereum:
		client, err = abi_client.NewEthereum(cfg)
		if err != nil {
			logger.Fatalf("Failed to instantiate eth abi client: %v", err)
		}
	case constants.Polygon:
		client, err = abi_client.NewPolygon(cfg)
		if err != nil {
			logger.Fatalf("Failed to instantiate polygon abi client: %v", err)
		}
	}
	return client
}
