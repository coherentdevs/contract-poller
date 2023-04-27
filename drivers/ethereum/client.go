package ethereum

import (
	"context"
	"fmt"
	"github.com/coherentopensource/chain-interactor/client/node"
	protos "github.com/coherentopensource/chain-interactor/protos/go/protos/chains/ethereum"
	"github.com/coherentopensource/contract-poller/abi_client"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/go-service-framework/util"
	"github.com/datadaodevs/go-service-framework/constants"
	"github.com/ethereum/go-ethereum/common"
	"github.com/metachris/eth-go-bindings/erc1155"
	"github.com/metachris/eth-go-bindings/erc165"
	"github.com/metachris/eth-go-bindings/erc20"
	"github.com/metachris/eth-go-bindings/erc721"
	"google.golang.org/protobuf/encoding/protojson"
	"strings"
)

const (
	owner = "0xe7A91167c495D881A58b56e780Bd6B1F51A3500e"
)

type client struct {
	logger     util.Logger
	blockchain constants.Blockchain
	node       node.Client
	abiSource  abi_client.Client
	cursor     uint64
}

// GetLatestBlockNumber gets the most recent block number
func (c *client) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	number, err := c.node.GetLatestBlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// GetBlockByNumber gets a block by number
func (c *client) GetBlockByNumber(ctx context.Context, blockNumber uint64) (*protos.Block, error) {
	res, err := c.node.GetBlockByNumber(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	data := &protos.Block{}
	if err := protojson.Unmarshal(res.Result, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (c *client) GetTracesForBlock(ctx context.Context, blockNumber uint64) ([]*protos.CallTrace, error) {
	// genesis block has no traces
	if blockNumber == 0 {
		return nil, nil
	}

	res, err := c.node.GetTracesForBlock(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	var rawTraces []*protos.CallTrace
	for _, trace := range res.Result {
		if trace.Error != nil {
			return nil, fmt.Errorf("%v", trace.Error)
		}
		rawTrace := &protos.CallTrace{}
		if err := protojson.Unmarshal(trace.Result, rawTrace); err != nil {
			return nil, err
		}
		rawTraces = append(rawTraces, rawTrace)
	}

	return rawTraces, nil
}

func (c *client) GetBlockReceipt(ctx context.Context, blockNumber uint64) ([]*protos.TransactionReceipt, error) {
	res, err := c.node.GetBlockReceipt(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	var rawReceipts []*protos.TransactionReceipt
	for _, receipt := range res.Result {
		rawReceipt := &protos.TransactionReceipt{}
		if err := protojson.Unmarshal(receipt, rawReceipt); err != nil {
			return nil, err
		}
		rawReceipts = append(rawReceipts, rawReceipt)
	}

	return rawReceipts, nil
}

func (c *client) GetContractMetadata(ctx context.Context, contractAddress string, blockNumber uint64) (*models.Contract, error) {
	contract := models.Contract{
		Address:    strings.ToLower(contractAddress),
		Blockchain: c.blockchain,
	}
	rpcClient := c.node.GetEthClient()

	// Validate that this is a contract
	resp, err := c.node.CodeAt(ctx, contractAddress, blockNumber)
	if err != nil {
		c.logger.Errorf("error getting code at address", err)
		return nil, err
	}

	if len(resp.Result) == 0 {
		contract.Standard = "user"
		return nil, nil
	}

	// Determine what type of contract it is
	address := common.HexToAddress(contractAddress)
	token, err := erc20.NewErc20(address, rpcClient)

	if err == nil {
		if isERC20(token) {
			// name, symbol, decimals are optional fields for ERC20
			contract.Standard = "erc20"
			name, _ := token.Name(nil)
			contract.Name = name
			symbol, _ := token.Symbol(nil)
			contract.Symbol = symbol
			decimals, err := token.Decimals(nil)
			if err != nil {
				return &contract, err
			}
			contract.Decimals = int32(decimals)
			return &contract, nil
		}
	}

	erc721Token, err := erc721.NewErc721(address, rpcClient)
	if err == nil {
		supportsInterface, err := erc721Token.SupportsInterface(nil, erc165.InterfaceIdErc721)
		if err == nil && supportsInterface {
			contract.Standard = "erc721"
			name, _ := erc721Token.Name(nil)
			contract.Name = name
			symbol, _ := erc721Token.Symbol(nil)
			contract.Symbol = symbol
			contract.Decimals = 0

			return &contract, nil
		}
	}

	erc1155Token, err := erc1155.NewErc1155(address, rpcClient)
	if err == nil {
		supportsInterface, err := erc1155Token.SupportsInterface(nil, erc165.InterfaceIdErc1155)
		if err == nil && supportsInterface {
			contract.Standard = "erc1155"
			return &contract, nil
		}
	}

	return &contract, nil
}

// Rough check if it's an ERC20 if it has some key reader interface methods
func isERC20(token *erc20.Erc20) bool {
	if _, err := token.BalanceOf(nil, common.HexToAddress(owner)); err != nil {
		return false
	} else if _, err := token.TotalSupply(nil); err != nil {
		return false
	} else if _, err := token.Allowance(nil, common.HexToAddress(owner), common.HexToAddress(owner)); err != nil {
		return false
	}
	return true
}
