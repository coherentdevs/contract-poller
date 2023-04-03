package node_client

import (
	"github.com/spf13/viper"

	"github.com/datadaodevs/go-service-framework/constants"
)

type Config struct {
	Blockchain constants.Blockchain

	EthNodeRPC  string
	PolyNodeRPC string
	OptNodeRPC  string
}

func NewConfig() *Config {
	setDefaults()

	viper.AutomaticEnv()
	config := Config{
		Blockchain:  constants.Blockchain(viper.GetString("blockchain")),
		EthNodeRPC:  viper.GetString("ethereum_node_rpc_endpoint"),
		OptNodeRPC:  viper.GetString("optimism_node_rpc_endpoint"),
		PolyNodeRPC: viper.GetString("polygon_node_rpc_endpoint"),
	}

	return &config
}

func setDefaults() {
	viper.SetDefault("blockchain", "ethereum")
}
