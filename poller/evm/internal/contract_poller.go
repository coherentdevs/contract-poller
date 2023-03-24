package contract_poller

import (
	"context"
	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/coherent-api/contract-poller/shared/service_framework"
	"github.com/nanmu42/etherscan-api"
	"github.com/pkg/errors"
	"strings"
)

var (
	ErrContractNotVerified = errors.New("contract source code not verified")
)

type contractPoller struct {
	config    *config.Config
	abiClient AbiClient
	evmClient EvmClient
	db        Database
	manager   *service_framework.Manager
}

type ContractPoller interface {
	Start(ctx context.Context) error
}

type AbiClient interface {
	ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error)
	GetContractABI(ctx context.Context, contractAddress string) (string, error)
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

type EvmClient interface {
	GetContract(contractAddress string) (*models.Contract, error)
}

func NewContractPoller(cfg *config.Config, db Database, abiClient AbiClient, evmClient EvmClient, manager *service_framework.Manager) (*contractPoller, error) {
	return &contractPoller{
		config:    cfg,
		abiClient: abiClient,
		evmClient: evmClient,
		db:        db,
		manager:   manager,
	}, nil
}

func (p *contractPoller) Start(ctx context.Context) error {
	return p.beginContractBackfiller(ctx)
}

func (p *contractPoller) beginContractBackfiller(ctx context.Context) error {
	contracts, err := p.db.GetContractsToBackfill()
	if err != nil {
		return err
	}
	updatedContracts := make([]models.Contract, 0)
	for _, contract := range contracts {
		contractMetadata, err := p.evmClient.GetContract(contract.Address)
		if err != nil {
			p.manager.Logger().Errorf("error from EVM Client: %v", err)
			continue
		}
		abiResp, err := p.abiClient.ContractSource(ctx, contract.Address, p.config.Blockchain)
		if err != nil {
			p.manager.Logger().Errorf("error from ABI Client: %v", err)
			continue
		}
		abi := ""
		officialName := ""
		if !errors.Is(err, ErrContractNotVerified) {
			officialName = abiResp.ContractName
		}
		if !(abiResp.ABI == "Contract source code not verified") && !(abiResp.ABI == "") {
			abi = abiResp.ABI
		}
		updatedContract := &models.Contract{
			Address:      strings.ToLower(contract.Address),
			Blockchain:   p.config.Blockchain,
			Name:         contractMetadata.Name,
			Symbol:       contractMetadata.Symbol,
			OfficialName: officialName,
			Standard:     contractMetadata.Standard,
			ABI:          abi,
			Decimals:     contractMetadata.Decimals,
		}
		updatedContracts = append(updatedContracts, *updatedContract)
	}
	backfillErr := p.db.UpdateContractsToBackfill(updatedContracts)
	numContracts, upsertErr := p.db.UpsertContracts(updatedContracts)
	if upsertErr != nil {
		return upsertErr
	}
	if backfillErr != nil {
		return backfillErr
	}

	p.manager.Logger().Infof("upserted %d contracts", numContracts)
	return nil
}
