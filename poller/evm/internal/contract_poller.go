package contract_poller

import (
	"context"
	"github.com/coherent-api/contract-poller/poller/evm/client/abi_client"
	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/db"
	"github.com/coherent-api/contract-poller/shared/go/service_framework"
)

type contractPoller struct {
	config          *config.Config
	etherscanClient abi_client.AbiClient
	db              db.Database
	manager         *service_framework.Manager
}

type ContractPoller interface {
	Start(ctx context.Context) error
}

func NewContractPoller(cfg *config.Config, db *db.DB, client *abi_client.RateLimitedClient, manager *service_framework.Manager) (*contractPoller, error) {

	return &contractPoller{
		config:          cfg,
		etherscanClient: client,
		db:              db,
		manager:         manager,
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
