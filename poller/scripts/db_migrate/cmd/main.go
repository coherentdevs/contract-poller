package main

import (
	"log"

	"github.com/coherent-api/contract-poller/poller/pkg/config"
	database "github.com/coherent-api/contract-poller/poller/pkg/db"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/service_framework"
)

func main() {
	manager, err := service_framework.NewManager()
	if err != nil {
		log.Fatalf("error creating new manager in db: %v", err)
	}

	cfg := config.NewConfig(manager)
	db := database.MustNewDB(cfg, manager)

	if err := db.Connection.AutoMigrate(
		&models.Contract{},
		&models.MethodFragment{},
		&models.EventFragment{},
	); err != nil {
		manager.Logger().Fatal(err)
	}

	manager.Logger().Infof("Done Migrating!")

	manager.WaitForInterrupt()
}