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
	TestContracts = []string{"0x00000000006c3852cbef3e08e8df289169ede581", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"}
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
	//Creates list of contracts we want to poll etherscan for
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
