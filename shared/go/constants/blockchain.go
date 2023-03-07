package constants

import (
	"github.com/coherent-api/contract-service/protos/go/protos/shared"
)

type Blockchain string

const (
	UNKNOWN             = "unknown"
	Ethereum Blockchain = "ethereum"
	Polygon  Blockchain = "polygon"
	Optimism Blockchain = "optimism"
	Goerli   Blockchain = "goerli"
)

func (b Blockchain) GetSymbol() string {
	switch b {
	case Ethereum:
		return NativeEthSymbol
	case Polygon:
		return NativePolygonSymbol
	case Optimism:
		return NativeOptimismSymbol
	case Goerli:
		return NativeGoerliSymbol
	}
	return UNKNOWN
}

var ProtoToBlockchain = map[shared.Blockchain]Blockchain{
	shared.Blockchain_OPTIMISM: Optimism,
	shared.Blockchain_POLYGON:  Polygon,
	shared.Blockchain_ETHEREUM: Ethereum,
	shared.Blockchain_GOERLI:   Goerli,
}

var BlockchainToProto = map[Blockchain]shared.Blockchain{
	Optimism: shared.Blockchain_OPTIMISM,
	Polygon:  shared.Blockchain_POLYGON,
	Ethereum: shared.Blockchain_ETHEREUM,
	Goerli:   shared.Blockchain_GOERLI,
}

type AddressType string

const (
	INVALID              = "invalid"
	CONTRACT AddressType = "contract"
	USER     AddressType = "user"
)
