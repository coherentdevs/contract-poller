package contract_poller

import (
	"context"

	"github.com/coherent-api/contract-poller/poller/evm/client/abi_client"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/coherent-api/contract-poller/shared/service_framework"
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

func MustNewABIClient(cfg *abi_client.Config, manager *service_framework.Manager) ABIClient {
	var client ABIClient
	var err error

	switch cfg.Blockchain {
	case constants.Ethereum:
		client, err = abi_client.NewEthereum(cfg)
		if err != nil {
			manager.Logger().With(err).Fatal("Failed to instantiate eth abi client")
		}
	case constants.Polygon:
		client, err = abi_client.NewPolygon(cfg)
		if err != nil {
			manager.Logger().With(err).Fatal("Failed to instantiate polygon abi client")
		}
	}
	return client
}
