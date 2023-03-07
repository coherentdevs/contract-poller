package config

import (
	"fmt"
	"time"

	"github.com/coherent-api/contract-service/shared/go/service_framework"
	"github.com/spf13/viper"
)

type Config struct {
	Host             string
	User             string
	Password         string
	DBName           string
	Port             int
	SSLMode          string
	CreateBatchSize  int
	ConnectionsLimit int
	QueryTimeout     time.Duration
	PeriodDuration   time.Duration

	manager *service_framework.Manager

	EtherscanAPIKey           string
	EtherscanRateMilliseconds time.Duration
	EtherscanRateRequests     int
	EtherscanErrorSleep       time.Duration
	EtherscanNetwork          etherscan.Network

	PolygonscanURL     string
	PolygonscanAPIKey  string
	PolygonscanTimeout time.Duration
}

func NewConfig(manager *service_framework.Manager) *Config {
	setDefaults()

	viper.AutomaticEnv()

	return &Config{
		Host:             viper.GetString("dbhost"),
		User:             viper.GetString("dbuser"),
		Password:         viper.GetString("dbpassword"),
		DBName:           viper.GetString("dbname"),
		Port:             viper.GetInt("dbport"),
		SSLMode:          viper.GetString("sslmode"),
		CreateBatchSize:  viper.GetInt("create_batch_size"),
		ConnectionsLimit: viper.GetInt("connections_limit"),
		QueryTimeout:     viper.GetDuration("query_timeout"),
		PeriodDuration:   viper.GetDuration("period_duration"),

		manager: manager,

		EtherscanAPIKey:           viper.GetString("etherscan_api_key"),
		EtherscanRateMilliseconds: viper.GetDuration("etherscan_rate_milliseconds"),
		EtherscanRateRequests:     viper.GetInt("etherscan_rate_requests"),
		EtherscanErrorSleep:       viper.GetDuration("etherscan_error_sleep"),
		EtherscanNetwork:          etherscan.Network(viper.GetString("etherscan_network")),

		PolygonscanURL:     viper.GetString("polygonscan_url"),
		PolygonscanAPIKey:  viper.GetString("polygonscan_api_key"),
		PolygonscanTimeout: viper.GetDuration("polygonscan_timeout"),
	}
}

func setDefaults() {
	viper.SetDefault("dbhost", "localhost")
	viper.SetDefault("dbpassword", "postgres")
	viper.SetDefault("dbuser", "postgres")
	viper.SetDefault("dbname", "db")
	viper.SetDefault("dbport", 5432)
	viper.SetDefault("sslmode", "disable")
	viper.SetDefault("create_batch_size", 10)
	viper.SetDefault("connections_limit", 1000)
	viper.SetDefault("query_timeout", "10s")
	viper.SetDefault("period_duration", "1s")
	viper.SetDefault("etherscan_api_key", "TB2F8U1GE54PA7EE32RUU6MW8Q8PIZNC25")
	viper.SetDefault("etherscan_rate_milliseconds", 100)
	viper.SetDefault("etherscan_rate_requests", 200)
	viper.SetDefault("etherscan_error_sleep", 1000)
	viper.SetDefault("polygonscan_url", "https://api.polygonscan.com/api?module=contract&action=getsourcecode")
	viper.SetDefault("polygonscan_api_key", "7C859GNQJRIZ2I5JAGSKNF839NWUH7RD34")
	viper.SetDefault("polygonscan_timeout", "10s")
	viper.SetDefault("etherscan_api_key", "XCX3K92XNGNJJRUGS8KBVVX76HJAHGZ57X")
}

func (c *Config) DSN() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.Port,
		c.SSLMode,
	)
	c.manager.Logger().Infof("DSN: %v", dsn)
	return dsn
}
