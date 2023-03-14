package config

import (
	"github.com/spf13/viper"
)

type Environment string

const (
	Local       Environment = "local"
	Test                    = "test"
	Development             = "development"
	Production              = "production"
)

type Config struct {
	AppName     string
	Env         Environment
	DatadogIP   string
	DatadogPort string
}

func NewConfig() *Config {
	setDefaults()

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	config := Config{
		AppName:     viper.GetString("app"),
		Env:         Environment(viper.GetString("env")),
		DatadogIP:   viper.GetString("datadog_ip"),
		DatadogPort: viper.GetString("datadog_port"),
	}

	return &config
}

func setDefaults() {
	viper.SetDefault("env", "local")
}
