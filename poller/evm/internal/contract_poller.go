package contract_poller

import (
	"context"
	"strings"

	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/datadaodevs/go-service-framework/constants"
	"github.com/datadaodevs/go-service-framework/util"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
)

var (
	ErrContractNotVerified = errors.New("contract source code not verified")
)

type contractPoller struct {
	blockchain  constants.Blockchain
	abiClient   ABIClient
	evmClient   EVMClient
	db          Database
	logger      util.Logger
	rateLimiter *rate.Limiter
}

type ContractPoller interface {
	Start(ctx context.Context) error
}

func NewContractPoller(blockchain constants.Blockchain, opts ...opt) *contractPoller {
	p := &contractPoller{
		blockchain: blockchain,
	}

	for _, fn := range opts {
		fn(p)
	}
	return p
}

func (p *contractPoller) Start(ctx context.Context) error {
	return p.beginContractBackfiller(ctx)
}

func (p *contractPoller) beginContractBackfiller(ctx context.Context) error {
	contracts, err := p.db.GetContractsToBackfill(p.blockchain)
	if err != nil {
		return err
	}
	updatedContracts := make([]models.Contract, 0)
	for _, contract := range contracts {

		if err := p.rateLimiter.Wait(ctx); err != nil {
			return err
		}

		contractMetadata, err := p.evmClient.GetContract(contract.Address)
		if err != nil {
			p.logger.Errorf("error from EVM Client: %v", err)
			continue
		}
		abiResp, err := p.abiClient.ContractSource(ctx, contract.Address, p.blockchain)
		if err != nil {
			p.logger.Errorf("error from ABI Client: %v", err)
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
			Blockchain:   p.blockchain,
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

	p.logger.Infof("upserted %d contracts", numContracts)
	return nil
}
