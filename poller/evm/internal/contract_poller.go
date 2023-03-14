package contract_poller

import (
	"context"
	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/coherent-api/contract-poller/shared/service_framework"
	"github.com/nanmu42/etherscan-api"
)

type contractPoller struct {
	config    *config.Config
	evmClient EvmClient
	db        Database
	manager   *service_framework.Manager
}

type ContractPoller interface {
	Start(ctx context.Context) error
}

type EvmClient interface {
	ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error)
	GetContractABI(ctx context.Context, contractAddress string) (string, error)
}

type Database interface {
	GetContractsToBackfill() ([]models.Contract, error)
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

func NewContractPoller(cfg *config.Config, db Database, client EvmClient, manager *service_framework.Manager) (*contractPoller, error) {
	return &contractPoller{
		config:    cfg,
		evmClient: client,
		db:        db,
		manager:   manager,
	}, nil
}

func (p *contractPoller) Start(ctx context.Context) error {
	return p.beginContractBackfiller(ctx)
}

func (p *contractPoller) beginContractBackfiller(ctx context.Context) error {
	//TODO: Implement this
	contracts, err := p.db.GetContractsToBackfill()
	if err != nil {
		return err
	}
	for _, contract := range contracts {
		//abi, err := p.etherscanClient.GetContractABI(ctx, contract.Address)
		//if err != nil {
		//	return err
		//}
		//contract := &models.Contract{
		//	Address: contract.Address,
		//	ABI:     abi,
		//}
		//err = p.db.UpdateContractByAddress(contract)
		//if err != nil {
		//	return err
		//}
		p.manager.Logger().Infof("Contract Address: %s", contract.Address)
	}
	return nil
}