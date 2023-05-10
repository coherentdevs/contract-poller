package optimism

import (
	"context"
	"fmt"
	protos "github.com/coherentopensource/chain-interactor/protos/go/protos/chains/optimism"
	"github.com/coherentopensource/go-service-framework/pool"
	"github.com/coherentopensource/go-service-framework/retry"
	"sync"
)

// FetchSequence defines the parallelizable steps in the fetch sequence
func (d *Driver) FetchSequence(blockHeight uint64) map[string]pool.Runner {
	d.client.cursor = blockHeight
	return map[string]pool.Runner{
		stageFetchReceipt: d.queueGetBlockAndTxReceiptByNumber(blockHeight),
		stageFetchTraces:  d.queueGetBlockTraceByNumber(blockHeight),
	}
}

// GetChainTipNumber gets the block number of the chaintip
func (d *Driver) GetChainTipNumber(ctx context.Context) (uint64, error) {
	var blockNum uint64
	var err error
	if err := retry.Exec(d.config.MaxRetries, func() error {
		blockNum, err = d.client.GetLatestBlockNumber(ctx)
		if err != nil {
			d.logger.Warnf("error thrown while trying to retrieve latest block number: %v", err)
			return err
		}
		return nil
	}, nil); err != nil {
		d.logger.Errorf("Max retries exceeded trying to get chaintip number: %v", err)
		return 0, err
	}

	return blockNum, nil
}

// getBlockTraceByNumber fetches all traces for a given block
func (d *Driver) getBlockTraceByNumber(ctx context.Context, blockHeight uint64) ([]*protos.CallTrace, error) {
	var traces []*protos.CallTrace
	var err error
	if err := retry.Exec(d.config.MaxRetries, func() error {
		traces, err = d.client.GetTracesForBlock(ctx, blockHeight)
		if err != nil {
			d.logger.Warnf("error thrown while trying to retrieve block trace: %d, %v", blockHeight, err)
			return err
		}

		return nil
	}, nil); err != nil {
		d.logger.Errorf("Max retries exceeded trying to get traces: %v", err)
		return nil, err
	}
	return traces, nil
}

func (d *Driver) getTransactionReceipt(ctx context.Context, txHash string) (*protos.TransactionReceipt, error) {
	var txReceipt *protos.TransactionReceipt
	var err error

	if err := retry.Exec(d.config.MaxRetries, func() error {
		txReceipt, err = d.client.GetTransactionReceipt(ctx, txHash)
		if err != nil {
			d.logger.Warnf("error thrown while trying to retrieve transaction receipt: %d, %v", txHash, err)
			return err
		}
		return nil
	}, nil); err != nil {
		d.logger.Errorf("max retries exceeded trying to get block by number: %v", err)
		return nil, err
	}

	return txReceipt, nil
}

// queueGetBlockTraceByNumber wraps GetBlockTraceByNumber in a queueable Runner func
func (d *Driver) queueGetBlockTraceByNumber(blockHeight uint64) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		return d.getBlockTraceByNumber(ctx, blockHeight)
	}
}

// getBlockByNumber fetches a full block by number
func (d *Driver) getBlockByNumber(ctx context.Context, blockHeight uint64) (*protos.Block, error) {
	var block *protos.Block
	var err error
	if err := retry.Exec(d.config.MaxRetries, func() error {
		block, err = d.client.GetBlockByNumber(ctx, blockHeight)
		if err != nil {
			d.logger.Warnf("error thrown while trying to retrieve block: %d, %v", blockHeight, err)
			return err
		}

		return nil
	}, nil); err != nil {
		d.logger.Errorf("max retries exceeded trying to get block by number: %v", err)
		return nil, err
	}

	return block, nil
}

// queueGetBlockAndTxReceiptByNumber wraps ReadBlockByNumber in a queueable Runner func
func (d *Driver) queueGetBlockAndTxReceiptByNumber(blockHeight uint64) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		block, err := d.getBlockByNumber(ctx, blockHeight)
		if err != nil {
			return nil, err
		}

		if len(block.Transactions) == 0 {
			return nil, fmt.Errorf("no transactions present in block %d", blockHeight)
		}

		receipts := make([]*protos.TransactionReceipt, len(block.Transactions))

		var wg sync.WaitGroup
		wg.Add(len(block.Transactions))
		for index, transaction := range block.Transactions {
			go func(ctx context.Context, i int, tx *protos.Transaction) {
				defer wg.Done()
				txReceipt, err := d.getTransactionReceipt(ctx, tx.Hash)
				if err != nil {
					d.logger.Errorf("error fetching transaction receipt with hash: %s, %v", tx.Hash, err)
				}
				receipts[i] = txReceipt
			}(ctx, index, transaction)

		}
		wg.Wait()

		return receipts, nil
	}
}
