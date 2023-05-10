package base

import (
	"context"
	"errors"
	"fmt"
	protos "github.com/coherentopensource/chain-interactor/protos/go/protos/chains/base"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/go-service-framework/pool"
	"github.com/coherentopensource/go-service-framework/retry"
	"strings"
	"sync"
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
		//stageFetchABI:      d.queueGetContractABI,
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
			statsdErr := d.metrics.Incr("etherscan_calls", []string{}, 1.0)
			if statsdErr != nil {
				d.logger.Warnf("error incrementing etherscan calls metric: %v", statsdErr)
			}
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
			d.logger.Errorf("Max retries exceeded trying to get contract metadata for %s: %v", address, err)
			return nil, err
		}

	}

	return contracts, nil
}

type callTraceNode struct {
	CallTrace       *protos.CallTrace
	Index           int64
	ContractAddress string
}

// flattenTraceAndReceipts flattens recursive trace calls & gets the contract address for each call (if it exists)
func (d *Driver) flattenTraceAndReceipts(receipts []*protos.TransactionReceipt, traces []*protos.CallTrace) ([]*callTraceNode, error) {

	if len(receipts) != len(traces) {
		return nil, errors.New(fmt.Sprintf("transactions and traces count don't match for block: %d %d != %d", d.client.cursor, len(receipts), len(traces)))
	}

	var bfsWG sync.WaitGroup
	var outputs []*callTraceNode
	mutex := sync.Mutex{}
	for i, callTrace := range traces {
		bfsWG.Add(1)
		go func(index int, callTrace *protos.CallTrace) {
			defer bfsWG.Done()

			queue := make([]*callTraceNode, 0)
			queue = append(
				queue,
				&callTraceNode{
					CallTrace: callTrace,
					Index:     int64(index),
				},
			)
			for len(queue) > 0 {
				currentNode := queue[0]
				queue = queue[1:]
				mutex.Lock()
				outputs = append(
					outputs, &callTraceNode{
						callTrace,
						currentNode.Index,
						receipts[index].GetContractAddress(),
					},
				)
				mutex.Unlock()
				for callIndex, call := range currentNode.CallTrace.Calls {
					queue = append(
						queue,
						&callTraceNode{
							CallTrace: call,
							Index:     int64(callIndex),
						},
					)
				}

			}
		}(i, callTrace)
	}
	bfsWG.Wait()

	return outputs, nil
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
	flattenedTraces, err := d.flattenTraceAndReceipts(receipts, traces)
	if err != nil {
		d.logger.Errorf("!!!!!!error flattening trace: %v", err)
	}
	for _, trace := range flattenedTraces {
		if trace.CallTrace.GetType() == "CREATE" && len(trace.ContractAddress) > 0 {
			deployments = append(deployments, trace.ContractAddress)
		}
	}
	err = d.metrics.Incr("contract_deployments", []string{}, float64(len(deployments)))
	if err != nil {
		d.logger.Warnf("error incrementing contract deployments metric: %v", err)
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
