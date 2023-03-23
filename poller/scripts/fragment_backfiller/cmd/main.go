package main

import (
	"log"

	"github.com/coherent-api/contract-poller/poller/pkg/config"
	database "github.com/coherent-api/contract-poller/poller/pkg/db"
	"github.com/coherent-api/contract-poller/shared/service_framework"
)

func main() {
	manager, err := service_framework.NewManager()
	if err != nil {
		log.Fatalf("error creating new manager in db: %v", err)
	}

	cfg := config.NewConfig(manager)
	db, err := database.NewDB(cfg, manager)
	if err != nil {
		manager.Logger().Fatal(err)
	}

	manager.PeriodicService(manager.Config.AppName, db.StartFragmentBackfiller, cfg.FragmentBuilderPeriodDuration)
	manager.WaitForInterrupt()
}
