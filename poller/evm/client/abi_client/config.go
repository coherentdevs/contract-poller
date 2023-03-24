package abi_client

import (
	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/nanmu42/etherscan-api"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	EtherscanAPIKey           string
	EtherscanRateMilliseconds time.Duration
	EtherscanRateRequests     int
	EtherscanErrorSleep       time.Duration
	EtherscanNetwork          etherscan.Network

	PolygonscanURL     string
	PolygonscanAPIKey  string
	PolygonscanTimeout time.Duration

	Blockchain constants.Blockchain
}

func NewConfig() *Config {
	setDefaults()

	viper.AutomaticEnv()
	return &Config{
		EtherscanAPIKey:           viper.GetString("etherscan_api_key"),
		EtherscanRateMilliseconds: viper.GetDuration("etherscan_rate_milliseconds"),
		EtherscanRateRequests:     viper.GetInt("etherscan_rate_requests"),
		EtherscanErrorSleep:       viper.GetDuration("etherscan_error_sleep"),
		EtherscanNetwork:          etherscan.Network(viper.GetString("etherscan_network")),

		PolygonscanURL:     viper.GetString("polygonscan_url"),
		PolygonscanAPIKey:  viper.GetString("polygonscan_api_key"),
		PolygonscanTimeout: viper.GetDuration("polygonscan_timeout"),

		Blockchain: constants.Blockchain(viper.GetString("blockchain")),
	}
}

func setDefaults() {
	viper.SetDefault("etherscan_rate_milliseconds", 100)
	viper.SetDefault("etherscan_rate_requests", 200)
	viper.SetDefault("etherscan_error_sleep", 1000)
	viper.SetDefault("etherscan_network", "api") //api - ethereum; api-optimistic - optimism
	viper.SetDefault("polygonscan_url", "https://api.polygonscan.com/api?module=contract&action=getsourcecode")
	viper.SetDefault("polygonscan_timeout", "10s")
	viper.SetDefault("blockchain", "ethereum")
}