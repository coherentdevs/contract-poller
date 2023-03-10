package main

import (
	"github.com/coherent-api/contract-poller/poller/evm/client/abi_client"
	contractPoller "github.com/coherent-api/contract-poller/poller/evm/internal"
	cfg "github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/db"
	"github.com/coherent-api/contract-poller/shared/go/service_framework"
)

func main() {
	manager, err := service_framework.NewManager()
	if err != nil {
		manager.Logger().Fatalf("error starting API manage: %v", err)
	}
	config := cfg.NewConfig(manager)
	abiClient := abi_client.NewClient(config)
	db, err := db.NewDB(config, manager)
	contractPoller, err := contractPoller.NewContractPoller(config, db, abiClient, manager)
	if err != nil {
		manager.Logger().Fatalf("could not initialize poller %v", err)
	}
	manager.PeriodicService(manager.Config.AppName, contractPoller.Start, config.PeriodDuration)
	manager.WaitForInterrupt()
}
