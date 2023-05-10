package optimism

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/go-service-framework/pool"
	"strings"
)

// Accumulate combines ABI and metadata to form complete contracts and fragments using result from the "fetch" step
func (d *Driver) Accumulate(res interface{}) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		resultSet, ok := res.(pool.ResultSet)
		if !ok {
			return nil, errors.New("no result set found")
		}
		contractMap, err := extractContractsWithMetadata(resultSet)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error extracting contracts: %s", err))
		}
		contracts := extractContracts(contractMap)
		methods, events, err := d.buildFragments(ctx, contracts)
		return models.ContractData{
			Contracts: contracts,
			Methods:   methods,
			Events:    events,
		}, nil
	}
}

// extractContractsWithABI extracts contract ABI from the generic pool.ResultSet from the fetch step
func extractContractsWithABI(set pool.ResultSet) (map[string]*models.Contract, error) {
	abiRes, ok := set[stageFetchABI]
	if !ok || abiRes == nil {
		return nil, errors.New("no abi for contracts")
	}
	bytes, _ := json.Marshal(abiRes)
	contracts := make(map[string]*models.Contract)
	err := json.Unmarshal(bytes, &contracts)
	if err != nil {
		return nil, err
	}
	return contracts, nil

}

// extractContractsWithMetadata extracts contract metadata from the generic pool.ResultSet from the fetch step
func extractContractsWithMetadata(set pool.ResultSet) (map[string]*models.Contract, error) {
	metadataRes, ok := set[stageFetchMetadata]
	if !ok || metadataRes == nil {
		return nil, errors.New("no metadata for contracts")
	}
	bytes, _ := json.Marshal(metadataRes)
	contracts := make(map[string]*models.Contract)
	err := json.Unmarshal(bytes, &contracts)
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

// buildFragments function to build method and event fragments from contracts using fragment builder
func (d *Driver) buildFragments(ctx context.Context, contracts []models.Contract) ([]models.MethodFragment, []models.EventFragment, error) {
	methods, events, err := d.fragmentBuilder.BuildFragmentsFromContracts(ctx, contracts)
	if err != nil {
		d.logger.Errorf("error building fragments: %v", err)
		return nil, nil, err
	}
	return methods, events, nil
}

// combineContracts helper function to combine ABI and metadata into a single contract using address mapping
func combineContracts(abis map[string]*models.Contract, metadata map[string]*models.Contract) []models.Contract {
	contractArr := make([]models.Contract, len(metadata))
	index := 0
	for address, contract := range abis {
		if metadataContract, ok := metadata[address]; ok {
			if contract != nil && metadataContract != nil {
				metadataContract.Address = address
				metadataContract.ABI = contract.ABI
				if metadata[address].Name == "" {
					metadataContract.Name = contract.Name
				}
				contractArr[index] = *metadata[address]
			}
		}
		index++
	}
	return contractArr
}

func extractContracts(contracts map[string]*models.Contract) []models.Contract {
	contractArr := make([]models.Contract, len(contracts))
	index := 0
	for address, contract := range contracts {
		if contract != nil {
			contractArr[index] = models.Contract{
				Address:      strings.ToLower(address),
				Blockchain:   contract.Blockchain,
				Name:         contract.Name,
				Symbol:       contract.Symbol,
				OfficialName: contract.OfficialName,
				Standard:     contract.Standard,
				ABI:          sanitizeString(contract.ABI),
				Decimals:     contract.Decimals,
			}
		}
		index++
	}
	return contractArr
}
