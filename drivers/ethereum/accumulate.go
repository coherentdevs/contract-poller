package ethereum

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/go-service-framework/pool"
)

// Accumulate combines ABI and metadata to form complete contracts and fragments using result from the "fetch" step
func (d *Driver) Accumulate(res interface{}) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		abis, err := extractContractsWithABI(res.(pool.ResultSet))
		if err != nil {
			d.logger.Errorf("error extracting ABI data: %s", err)
		}
		metadata, err := extractContractsWithMetadata(res.(pool.ResultSet))
		if err != nil {
			d.logger.Errorf("error extracting metadata: %s", err)
		}
		contracts := combineContracts(abis, metadata)
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
