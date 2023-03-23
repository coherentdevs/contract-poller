package db

import (
	"fmt"
	"github.com/coherent-api/contract-poller/shared/constants"
	"time"

	"github.com/coherent-api/contract-poller/shared/service_framework"
	"github.com/spf13/viper"
)

type Config struct {
	Host                          string
	User                          string
	Password                      string
	DBName                        string
	Port                          int
	SSLMode                       string
	CreateBatchSize               int
	ConnectionsLimit              int
	QueryTimeout                  time.Duration
	PeriodDuration                time.Duration
	FragmentBuilderPeriodDuration time.Duration
	FragmentBatchSize             int

	manager *service_framework.Manager

	Blockchain constants.Blockchain
}

func NewConfig(manager *service_framework.Manager) *Config {
	setDefaults()

	viper.AutomaticEnv()

	return &Config{
		Host:                          viper.GetString("dbhost"),
		User:                          viper.GetString("dbuser"),
		Password:                      viper.GetString("dbpassword"),
		DBName:                        viper.GetString("dbname"),
		Port:                          viper.GetInt("dbport"),
		SSLMode:                       viper.GetString("sslmode"),
		CreateBatchSize:               viper.GetInt("create_batch_size"),
		ConnectionsLimit:              viper.GetInt("connections_limit"),
		QueryTimeout:                  viper.GetDuration("query_timeout"),
		PeriodDuration:                viper.GetDuration("period_duration"),
		FragmentBuilderPeriodDuration: viper.GetDuration("fragment_builder_period_duration"),
		FragmentBatchSize:             viper.GetInt("fragment_batch_size"),

		manager: manager,

		Blockchain: constants.Blockchain(viper.GetString("blockchain")),
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
	viper.SetDefault("period_duration", "10s")
	viper.SetDefault("fragment_builder_period_duration", "1000000s")
	viper.SetDefault("fragment_batch_size", 10000)
	viper.SetDefault("blockchain", "ethereum")
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
