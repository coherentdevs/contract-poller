package node_client

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/metachris/eth-go-bindings/erc1155"
	"github.com/metachris/eth-go-bindings/erc165"
	"github.com/metachris/eth-go-bindings/erc20"
	"github.com/metachris/eth-go-bindings/erc721"

	"strings"
)

func (c *evmClient) GetContractFromNode(ctx context.Context, req *GetContractRequest) (*GetContractResponse, error) {
	resp := &GetContractResponse{Address: req.Address}

	address := common.HexToAddress(req.Address)

	// Validate that this is a contract
	bytes, err := c.parsedClient.CodeAt(ctx, address, nil)
	if err != nil {
		return nil, err
	}
	if len(bytes) == 0 {
		resp.Type = AddressType_USER
		return resp, nil
	}

	resp.Bytecode = bytes

	// Determine what type of contract it is
	token, err := erc20.NewErc20(address, c.parsedClient)
	if err == nil {
		if isERC20(token) {
			// name, symbol, decimals are optional fields for ERC20
			resp.Type = AddressType_ERC20
			name, err := token.Name(nil)
			if err != nil {
				return nil, err
			}
			resp.Name = name
			symbol, err := token.Symbol(nil)
			if err != nil {
				return nil, err
			}
			resp.Symbol = symbol
			decimals, err := token.Decimals(nil)
			if err != nil {
				return nil, err
			}
			resp.Decimals = int32(decimals)
			resp.Name = sanitizeString(resp.Name)
			resp.Symbol = sanitizeString(resp.Symbol)
		} else {
			resp.Type = AddressType_CONTRACT
		}
	}

	erc721Token, err := erc721.NewErc721(address, c.parsedClient)
	if err == nil {
		supportsInterface, err := erc721Token.SupportsInterface(nil, erc165.InterfaceIdErc721)
		if err == nil && supportsInterface {
			resp.Type = AddressType_ERC721
			name, err := erc721Token.Name(nil)
			if err != nil {
				return nil, err
			}
			resp.Name = name
			symbol, err := erc721Token.Symbol(nil)
			if err != nil {
				return nil, err
			}
			resp.Symbol = symbol
			resp.Decimals = 0

			resp.Name = sanitizeString(resp.Name)
			resp.Symbol = sanitizeString(resp.Symbol)

			return resp, nil
		}
	}

	erc1155Token, err := erc1155.NewErc1155(address, c.parsedClient)
	if err == nil {
		supportsInterface, err := erc1155Token.SupportsInterface(nil, erc165.InterfaceIdErc1155)
		if err == nil && supportsInterface {
			resp.Type = AddressType_ERC1155
			return resp, nil
		}
	} else if err != nil {
		resp.Type = AddressType_CONTRACT
		return resp, err
	}

	return resp, nil
}

// Rough check if it's an ERC20 if it has some key reader interface methods
func isERC20(token *erc20.Erc20) bool {
	if _, err := token.BalanceOf(nil, common.HexToAddress("0xe7A91167c495D881A58b56e780Bd6B1F51A3500e")); err != nil {
		return false
	} else if _, err := token.TotalSupply(nil); err != nil {
		return false
	} else if _, err := token.Allowance(nil, common.HexToAddress("0xe7A91167c495D881A58b56e780Bd6B1F51A3500e"), common.HexToAddress("0xe7A91167c495D881A58b56e780Bd6B1F51A3500e")); err != nil {
		return false
	}
	return true
}

func sanitizeString(str string) string {
	return strings.ToValidUTF8(strings.ReplaceAll(str, "\x00", ""), "")
}
