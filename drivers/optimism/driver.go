package optimism

import (
	"github.com/caarlos0/env/v7"
	nodeClient "github.com/coherentopensource/chain-interactor/client/node"
	"github.com/coherentopensource/contract-poller/abi_client"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/contract-poller/util/fragment_builder"
	"github.com/coherentopensource/go-service-framework/constants"
	"github.com/coherentopensource/go-service-framework/database"
	"github.com/coherentopensource/go-service-framework/database/postgres"
	"github.com/coherentopensource/go-service-framework/util"
)

const (
	stageFetchReceipt   = "fetch.block"
	stageFetchTraces    = "fetch.traces"
	stageFetchMetadata  = "fetch.metadata"
	stageFetchABI       = "fetch.abi"
	stageFetchAddresses = "fetch.addresses"
)

// Driver is the container for all ETL business logic
type Driver struct {
	client          *client
	logger          util.Logger
	metrics         util.Metrics
	config          *Config
	database        *database.Database
	fragmentBuilder *fragment_builder.Constructor
}

// NewDriver constructs a new Driver
func NewDriver(cfg *Config, nodeClient nodeClient.Client, logger util.Logger, metrics util.Metrics, cursor uint64) (*Driver, error) {
	abiClient, err := abi_client.NewEthereumABIClient(abi_client.MustParseConfig(logger))
	if err != nil {
		return nil, err
	}
	driver := &Driver{
		client:          &client{node: nodeClient, abiSource: abiClient, logger: logger, cursor: cursor, blockchain: cfg.Blockchain},
		logger:          logger,
		metrics:         metrics,
		config:          cfg,
		fragmentBuilder: &fragment_builder.Constructor{},
	}
	var postgresCfg database.Config
	if err := env.Parse(&postgresCfg); err != nil {
		return nil, err
	}
	db, err := postgres.NewPostgresDB(driver, &postgresCfg, logger)
	if err != nil {
		return nil, err
	}
	err = db.Migrate(&models.Contract{}, &models.EventFragment{}, &models.MethodFragment{})
	if err != nil {
		return nil, err
	}
	driver.database = db
	return driver, nil
}

// Blockchain returns the name of the blockchain
func (d *Driver) Blockchain() string {
	return string(constants.Optimism)
}
