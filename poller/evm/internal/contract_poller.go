package contract_poller

import (
	"context"
	"github.com/coherent-api/contract-service/poller/evm/client/abi_client"
	"github.com/coherent-api/contract-service/poller/pkg/config"
	"github.com/coherent-api/contract-service/poller/pkg/db"
	"github.com/coherent-api/contract-service/poller/pkg/models"
	"github.com/coherent-api/contract-service/shared/go/service_framework"
)

type contractPoller struct {
	config          *config.Config
	etherscanClient *abi_client.RateLimitedClient
	db              *db.DB
	manager         *service_framework.Manager
}

type ContractPoller interface {
	Start(ctx context.Context) error
}

func NewContractPoller(cfg *config.Config, manager *service_framework.Manager) (*contractPoller, error) {
	etherscanClient := abi_client.NewClient(cfg)
	db, err := db.NewDB(cfg, manager)
	if err != nil {
		return nil, err
	}
	return &contractPoller{
		config:          cfg,
		etherscanClient: etherscanClient,
		db:              db,
		manager:         manager,
	}, nil
}

func (p *contractPoller) Start(ctx context.Context) error {
	return p.beginContractBackfiller(ctx)
}

func (p *contractPoller) beginContractBackfiller(ctx context.Context) error {
	//TODO: Implement this
	for _, contract := range p.db.Contracts {
		abi, err := p.etherscanClient.GetContractABI(ctx, contract.Address)
		if err != nil {
			return err
		}
		contract := &models.Contract{
			Address: contract.Address,
			ABI:     abi,
		}
		err = p.db.UpdateContractByAddress(contract)
		if err != nil {
			return err
		}
	}
	return nil
}
