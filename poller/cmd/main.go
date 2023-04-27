package main

import (
	"fmt"
	nodeClient "github.com/coherentopensource/chain-interactor/client/node"
	"github.com/coherentopensource/contract-poller/poller/internal/chains"
	svc_config "github.com/coherentopensource/contract-poller/poller/internal/config"
	"github.com/coherentopensource/contract-poller/poller/pkg/cache"
	"github.com/coherentopensource/contract-poller/poller/server"
	"github.com/coherentopensource/go-service-framework/constants"
	poller "github.com/coherentopensource/go-service-framework/contract_poller"
	"github.com/coherentopensource/go-service-framework/manager"
	"github.com/coherentopensource/go-service-framework/pool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/pprof"
)

func main() {
	mgr := manager.New(manager.WithoutGracefulShutdown())

	//	If local, force read from .env file
	if mgr.Env() == manager.EnvLocal {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	cfg := svc_config.MustNewConfig(mgr.Logger())

	nodeClient := nodeClient.MustNewClient(&cfg.Node, mgr.Logger())
	cache := cache.NewCache(&cfg.Cache, mgr.Logger())
	tt := pool.NewThrottler(cfg.FetcherPoolThrottleBandwidth, cfg.FetcherPoolThrottleDuration)
	wp1 := pool.NewWorkerPool(
		"fetchSequence",
		pool.WithOutputChannel(),
		pool.WithThrottler(tt),
		pool.WithLogger(mgr.Logger()),
		pool.WithBandwidth(cfg.FetcherPoolBandwidth),
	)
	wp2 := pool.NewWorkerPool(
		"fetchers",
		pool.WithOutputChannel(),
		pool.WithThrottler(tt),
		pool.WithLogger(mgr.Logger()),
		pool.WithBandwidth(cfg.FetcherPoolBandwidth),
	)

	wp3 := pool.NewWorkerPool(
		"accumulator",
		pool.WithOutputChannel(),
		pool.WithLogger(mgr.Logger()),
		pool.WithBandwidth(cfg.AccumulatorPoolBandwidth),
	)
	wp4 := pool.NewWorkerPool(
		"writer",
		pool.WithLogger(mgr.Logger()),
		pool.WithBandwidth(cfg.WriterPoolBandwidth),
	)
	cursor, err := cache.GetCurrentBlockNumber(mgr.Context(), fmt.Sprintf("contract_poller-%s-%s", cfg.Node.Blockchain, constants.BlockKey))
	if err != nil {
		mgr.Logger().Warnf("could not get current block number from cache: %v", err)
	}
	driver := chains.MustInitializeDriver(cfg.Poller.Blockchain, nodeClient, mgr.Logger(), cursor)

	p := poller.New(
		&cfg.Poller,
		driver,
		poller.WithCache(cache),
		poller.WithAddressFetchPool(wp1),
		poller.WithFetchPool(wp2),
		poller.WithAccumulatePool(wp3),
		poller.WithWritePool(wp4),
		poller.WithLogger(mgr.Logger()),
		poller.WithMetrics(mgr.Metrics()),
	)

	srv := server.New(p, mgr.Logger())
	apiSrv := http.Server{
		Addr:    cfg.HttpServePort,
		Handler: srv.Router(),
	}

	//	Background Services
	mgr.RegisterBackgroundSvc("fetch address pool", wp1.Start, wp1.Stop)
	mgr.RegisterBackgroundSvc("fetcher pool", wp2.Start, wp2.Stop)
	mgr.RegisterBackgroundSvc("accumulator pool", wp3.Start, wp3.Stop)
	mgr.RegisterBackgroundSvc("writer pool", wp4.Start, wp4.Stop)
	mgr.RegisterBackgroundSvc("poller", p.Start, p.Stop)
	mgr.RegisterBackgroundSvc("throttler", tt.Start, tt.Stop)

	//	HTTP server
	mgr.RegisterHttpServer("api", &apiSrv)
	mgr.RegisterHttpServer("pprof", getPprofServer(cfg))

	mgr.WaitForInterrupt()
}

func getPprofServer(cfg *svc_config.Config) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return &http.Server{
		Addr:    cfg.PprofServePort,
		Handler: mux,
	}
}
