package db

import (
	"context"
	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/service_framework"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

var (
	TestContracts = []string{"0x00000000006c3852cbef3e08e8df289169ede581", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"}
)

type DB struct {
	Connection *gorm.DB
	Config     *config.Config
	manager    *service_framework.Manager
}

func NewDB(cfg *config.Config, manager *service_framework.Manager) (*DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Silent),
		CreateBatchSize: cfg.CreateBatchSize,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	manager.Logger().Infof("db initialized with connection limit: %d", cfg.ConnectionsLimit)
	sqlDB.SetMaxOpenConns(cfg.ConnectionsLimit)
	sqlDB.SetMaxIdleConns(cfg.ConnectionsLimit)
	sqlDB.SetConnMaxLifetime(time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Minute)

	return &DB{
		Connection: db,
		Config:     cfg,
		manager:    manager,
	}, nil
}

func MustNewDB(cfg *config.Config, manager *service_framework.Manager) *DB {
	db, err := NewDB(cfg, manager)
	if err != nil {
		manager.Logger().Fatalf("failed to initialize db: %v", err)
	}
	return db
}

func (db *DB) GetContractsToBackfill() ([]models.Contract, error) {
	//TODO: Temporary solution for local development. This should be replaced with a query to the database
	contractList := make([]models.Contract, 0)
	for _, contractAddress := range TestContracts {
		contractList = append(contractList, models.Contract{Address: contractAddress})
	}
	return contractList, nil
}

func (db *DB) EmitQueryMetric(err error, query string) error {
	if err != nil {
		db.manager.Logger().Errorf("Query Timeout: %s with error: %v", query, err)
		return err
	}
	return nil
}

func (db *DB) SanitizeString(str string) string {
	return strings.ToValidUTF8(strings.ReplaceAll(str, "\x00", ""), "")
}

func (db *DB) StartFragmentBackfiller(ctx context.Context) error {
	return db.BuildFragmentsFromContracts(ctx)
}

func (db *DB) UpdateContractsToBackfill(updatedContracts []models.Contract) error {
	//TODO: This is used for local development. GetContractsToBackfill should handle this
	contractsToBackfill := make([]string, 0)
	for _, contract := range updatedContracts {
		if !contains(TestContracts, contract.Address) {
			contractsToBackfill = append(contractsToBackfill, contract.Address)
		}
	}
	TestContracts = contractsToBackfill
	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
