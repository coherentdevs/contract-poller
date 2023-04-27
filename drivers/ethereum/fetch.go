package ethereum

import (
	"context"
	"errors"
	protos "github.com/coherentopensource/chain-interactor/protos/go/protos/chains/ethereum"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/go-service-framework/pool"
	"github.com/coherentopensource/go-service-framework/retry"
	"strings"
)

func (d *Driver) queueGetContractABI(res interface{}) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		deployments := d.construct(res)
		return d.FetchABI(ctx, deployments)
	}
}

func (d *Driver) queueGetContractMetadata(res interface{}) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		deployments := d.construct(res)
		return d.FetchMetadata(ctx, deployments, d.client.cursor)
	}
}
func (d *Driver) Fetchers() map[string]pool.FeedTransformer {
	return map[string]pool.FeedTransformer{
		stageFetchABI:      d.queueGetContractABI,
		stageFetchMetadata: d.queueGetContractMetadata,
	}
}

func (d *Driver) FetchABI(ctx context.Context, deployments []string) (map[string]*models.Contract, error) {
	if len(deployments) == 0 {
		return nil, nil
	}
	contractABIs := make(map[string]*models.Contract, 0)
	for _, address := range deployments {
		if err := retry.Exec(d.config.MaxRetries, func() error {
			resp, err := d.client.abiSource.ContractSource(ctx, address)
			if err != nil {
				d.logger.Warnf("error thrown while trying to retrieve abi for contract: %v", err)
			}
			if resp != nil {
				contractABIs[address] = &models.Contract{
					Address:    strings.ToLower(address),
					Name:       resp.ContractName,
					Blockchain: d.config.Blockchain,
				}
				if resp.ABI != "Contract source code not verified" {
					contractABIs[address].ABI = sanitizeString(resp.ABI)
				}

				return nil
			}
			return errors.New("no abi found")
		}, nil); err != nil {
			d.logger.Errorf("Max retries exceeded trying to get chaintip number: %v", err)
			return nil, err
		}

	}

	return contractABIs, nil
}

func (d *Driver) FetchMetadata(ctx context.Context, deployments []string, blockNumber uint64) (map[string]*models.Contract, error) {
	if len(deployments) == 0 {
		return nil, nil
	}
	contracts := make(map[string]*models.Contract)
	for _, address := range deployments {
		if err := retry.Exec(d.config.MaxRetries, func() error {
			contract, err := d.client.GetContractMetadata(ctx, address, blockNumber)
			if err != nil {
				d.logger.Warnf("error thrown while trying to retrieve metadata for contract %s: %v", address, err)
				return err
			}
			contracts[address] = contract
			return nil
		}, nil); err != nil {
			d.logger.Errorf("Max retries exceeded trying to get chaintip number: %v", err)
			return nil, err
		}

	}

	return contracts, nil
}

func (d *Driver) construct(res interface{}) []string {
	set, ok := res.(pool.ResultSet)
	if !ok {
		d.logger.Warn("result is not expected type")
		return nil
	}
	receipts, err := d.extractReceipts(set)
	if err != nil {
		return nil
	}
	traces, err := d.extractTraces(set)
	if err != nil {
		return nil
	}

	var deployments []string
	for index, trace := range traces {
		receipt := receipts[index]
		if trace.GetType() == "CREATE" || trace.GetType() == "CREATE2" {
			deployments = append(deployments, receipt.GetContractAddress())
		}
	}
	return deployments
}

// extractReceipts extracts receipts from the generic ResultSet from the fetch step
func (d *Driver) extractReceipts(set pool.ResultSet) ([]*protos.TransactionReceipt, error) {
	receiptsRes, ok := set[stageFetchReceipt]
	if !ok {
		return nil, errors.New("no receipts data")
	}
	receipts, ok := receiptsRes.([]*protos.TransactionReceipt)
	if !ok {
		return nil, errors.New("incorrect data type for transaction receipts")
	}
	return receipts, nil
}

// extractTraces extracts traces from the generic ResultSet from the fetch step
func (d *Driver) extractTraces(set pool.ResultSet) ([]*protos.CallTrace, error) {
	tracesRes, ok := set[stageFetchTraces]
	if !ok {
		return nil, errors.New("no traces data")
	}

	traces, ok := tracesRes.([]*protos.CallTrace)
	if !ok {
		return nil, errors.New("incorrect data type for traces")
	}

	return traces, nil
}

func sanitizeString(str string) string {
	return strings.ToValidUTF8(strings.ReplaceAll(str, "\x00", ""), "")
}
