package config

import (
	"github.com/coherent-api/contract-poller/shared/service_framework"
	"github.com/datadaodevs/go-service-framework/constants"
	"github.com/spf13/viper"
)

type Config struct {
	manager    *service_framework.Manager
	Blockchain constants.Blockchain
}

func NewConfig(manager *service_framework.Manager) *Config {
	setDefaults()

	viper.AutomaticEnv()

	return &Config{
		manager: manager,

		Blockchain: constants.Blockchain(viper.GetString("blockchain")),
	}
}

func setDefaults() {
	viper.SetDefault("blockchain", "ethereum")
}
