package main

import (
	"time"

	"github.com/coherent-api/contract-poller/poller/evm/client/abi_client"
	node_client "github.com/coherent-api/contract-poller/poller/evm/client/node_client"
	contractPoller "github.com/coherent-api/contract-poller/poller/evm/internal"
	cfg "github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/poller/pkg/db"
	"github.com/coherent-api/contract-poller/shared/service_framework"
	"golang.org/x/time/rate"
)

func main() {
	manager, err := service_framework.NewManager()
	if err != nil {
		manager.Logger().Fatalf("error starting API manage: %v", err)
	}
	dbConfig := db.NewConfig(manager)
	config := cfg.NewConfig(manager)
	abiCfg := abi_client.NewConfig()
	abiClient := contractPoller.MustNewABIClient(abiCfg.Blockchain, abiCfg, manager.Logger())
	evmCfg := node_client.NewConfig()
	evmClient := node_client.MustNewClient(evmCfg, manager)
	db := db.MustNewDB(dbConfig, manager)
	rateLimiter := rate.NewLimiter(rate.Every(abiCfg.AbiClientRateMilliseconds*time.Millisecond), abiCfg.AbiClientRateRequests)

	contractPoller := contractPoller.NewContractPoller(
		config.Blockchain,
		contractPoller.WithABIClient(abiClient),
		contractPoller.WithDatabase(db),
		contractPoller.WithNodeClient(evmClient),
		contractPoller.WithRateLimiter(rateLimiter),
		contractPoller.WithLogger(manager.Logger()),
	)

	if err != nil {
		manager.Logger().Fatalf("could not initialize poller %v", err)
	}
	manager.PeriodicService(manager.Config.AppName, contractPoller.Start, dbConfig.PeriodDuration)
	manager.WaitForInterrupt()
}
