package main

import (
	"github.com/coherent-api/contract-poller/poller/evm/client/abi_client"
	"github.com/coherent-api/contract-poller/poller/evm/client/evm_client"
	contractPoller "github.com/coherent-api/contract-poller/poller/evm/internal"
	cfg "github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/db"
	"github.com/coherent-api/contract-poller/shared/service_framework"
)

func main() {
	manager, err := service_framework.NewManager()
	if err != nil {
		manager.Logger().Fatalf("error starting API manage: %v", err)
	}
	config := cfg.NewConfig(manager)
	abiCfg := abi_client.NewConfig()
	abiClient := abi_client.NewClient(abiCfg)
	evmCfg := evm_client.NewConfig()
	evmClient := evm_client.MustNewClient(evmCfg, manager)
	db := db.MustNewDB(config, manager)
	contractPoller, err := contractPoller.NewContractPoller(config, db, abiClient, evmClient, manager)
	if err != nil {
		manager.Logger().Fatalf("could not initialize poller %v", err)
	}
	manager.PeriodicService(manager.Config.AppName, contractPoller.Start, config.PeriodDuration)
	manager.WaitForInterrupt()
}
