package constants

var WETHMapping = map[Blockchain]string{
	Ethereum: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	Polygon:  "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619",
	Optimism: "0x4200000000000000000000000000000000000006",
}

var USDCMapping = map[Blockchain]string{
	Ethereum: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	Polygon:  "0x2791bca1f2de4661ed88a30c99a7a9449aa84174",
	Optimism: "0x7f5c764cbc14f9669b88837ca1490cca17c31607",
}

const (
	ERC20           = "ERC20"
	ERC721          = "ERC721"
	ERC1155         = "ERC1155"
	Native          = "native"
	ERC20Decimals   = 18
	ERC20DecimalStr = "18"
	User            = "user"
	Address         = "address"

	NativeEthName        = "ETH"
	NativeEthSymbol      = "ETH"
	NativeEthType        = "ETH"
	NativePolygonSymbol  = "MATIC"
	NativeOptimismSymbol = "ETH"
	NativeToken          = "native_token"
	NativeTransfer       = "transfer"

	MintToken = "mint"
	MintNFT   = "mint"
	BurnToken = "burn"
	BurnNFT   = "burn"

	NFTAsset    = "nft"
	NFTTransfer = "transfer"

	TokenTransfer = "transfer"
	TokenAsset    = "token"
	USD           = "USD"
	WETH          = "weth"
	One           = "1.0"
	Transfer      = "Transfer"
)
