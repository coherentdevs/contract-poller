package main

import (
	contractPoller "github.com/coherent-api/contract-poller/poller/evm/internal"
	contractPollerCfg "github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/shared/go/service_framework"
	"time"
)

func main() {
	manager, err := service_framework.NewManager()
	if err != nil {
		manager.Logger().Fatalf("error starting API manage: %v", err)
	}

	config := contractPollerCfg.NewConfig(manager)
	contractPoller, err := contractPoller.NewContractPoller(config, manager)
	if err != nil {
		manager.Logger().Fatalf("could not initialize poller %v", err)
	}

	manager.PeriodicService(manager.Config.AppName, contractPoller.Start, config.PeriodDuration*time.Second)
	manager.WaitForInterrupt()
}
