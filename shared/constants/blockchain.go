package constants

import (
	"github.com/coherentopensource/contract-poller/protos/go/protos/shared"
)

type Blockchain string

const (
	UNKNOWN             = "unknown"
	Ethereum Blockchain = "ethereum"
	Polygon  Blockchain = "polygon"
	Optimism Blockchain = "optimism"
)

func (b Blockchain) GetSymbol() string {
	switch b {
	case Ethereum:
		return NativeEthSymbol
	case Polygon:
		return NativePolygonSymbol
	case Optimism:
		return NativeOptimismSymbol
	}
	return UNKNOWN
}

var ProtoToBlockchain = map[shared.Blockchain]Blockchain{
	shared.Blockchain_OPTIMISM: Optimism,
	shared.Blockchain_POLYGON:  Polygon,
	shared.Blockchain_ETHEREUM: Ethereum,
}

var BlockchainToProto = map[Blockchain]shared.Blockchain{
	Optimism: shared.Blockchain_OPTIMISM,
	Polygon:  shared.Blockchain_POLYGON,
	Ethereum: shared.Blockchain_ETHEREUM,
}

type AddressType string

const (
	INVALID              = "invalid"
	CONTRACT AddressType = "contract"
	USER     AddressType = "user"
)
