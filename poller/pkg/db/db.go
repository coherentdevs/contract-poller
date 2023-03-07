package db

import (
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"strings"
	"time"

	"google.golang.org/api/option"
	"gorm.io/gorm/logger"

	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/shared/go/service_framework"

	"cloud.google.com/go/storage"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB

	manager   *service_framework.Manager
	gcsClient *storage.Client

	Contracts []models.Contract
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
	client, err := storage.NewClient(manager.Context(), option.WithoutAuthentication())
	return &DB{
		Connection: db,
		manager:    manager,
		gcsClient:  client,
	}, nil
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
