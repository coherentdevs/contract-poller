package evm_client

import (
	"time"

	"github.com/spf13/viper"

	"github.com/coherent-api/contract-poller/shared/constants"
)

type Config struct {
	Blockchain constants.Blockchain

	EthNodeRPC                string
	PolyNodeRPC               string
	GoerliNodeRPC             string
	OptNodeRPC                string
	EnrichTransactionsTimeout time.Duration

	FetchBlockTimeout time.Duration
}

func NewConfig() *Config {
	setDefaults()

	viper.AutomaticEnv()
	config := Config{
		Blockchain:                constants.Blockchain(viper.GetString("blockchain")),
		EthNodeRPC:                viper.GetString("ethereum_node_rpc_endpoint"),
		OptNodeRPC:                viper.GetString("optimism_node_rpc_endpoint"),
		PolyNodeRPC:               viper.GetString("polygon_node_rpc_endpoint"),
		GoerliNodeRPC:             viper.GetString("goerli_node_rpc_endpoint"),
		EnrichTransactionsTimeout: viper.GetDuration("enrich_transactions_timeout"),
		FetchBlockTimeout:         viper.GetDuration("fetch_block_timeout"),
	}

	return &config
}

func setDefaults() {
	viper.SetDefault("rpc_retries", 2)
	viper.SetDefault("rpc_timeout", "20000ms")
	viper.SetDefault("fetch_block_timeout", "14s")
	viper.SetDefault("blockchain", "ethereum")
}
