package db

import (
	"cloud.google.com/go/storage"
	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/go/service_framework"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

var (
	TestContracts = []string{"0x00000000006c3852cbef3e08e8df289169ede581", "0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45", "0xdac17f958d2ee523a2206206994597c13d831ec7", "0x7a250d5630b4cf539739df2c5dacb4c659f2488d", "0xef1c6e67703c7bd7107eed8303fbe6ec2554bf6b", "0x06450dee7fd2fb8e39061434babcfc05599a6fb8", "0x000000000000ad05ccc4f10045630fb830b95127", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0x881d40237659c251811cec9c364ef91dc08d300c", "0x0de8bf93da2f7eecb3d9169422413a9bef4ef628", "0x1111111254fb6c44bac0bed2854e76f90643097d", "0x83c8f28c26bf6aaca652df1dbbe0e1b56f8baba2", "0x283af0b28c62c092c9727f1ee09c02ca627eb7f5", "0x5e4e65926ba27467555eb562121fac00d24e9dd2", "0x1c479675ad559dc151f6ec7ed3fbf8cee79582b6", "0xa9d1e08c7793af67e9d92fe308d5697fb81d3e43", "0xe66b31678d6c16e9ebf358268a790b763c133750", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0x0a252663dbcc0b073063d6420a40319e438cfa59", "0x44e94034afce2dd3cd5eb62528f239686fc8f162", "0x1111111254eeb25477b68fb85ed929f73a960582", "0xc36442b4a4522e871399cd717abdd847ab11fe88", "0x39da41747a83aee658334415666f3ef92dd0d541"}
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
	if err != nil {
		return nil, err
	}
	//Creates list of contracts we want to poll etherscan for
	contractList := make([]models.Contract, 0)
	for _, contractAddress := range TestContracts {
		contractList = append(contractList, models.Contract{Address: contractAddress})
	}
	return &DB{
		Connection: db,
		manager:    manager,
		Contracts:  contractList,
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
