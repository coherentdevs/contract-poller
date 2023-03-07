package service_framework

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"time"

	"go.uber.org/zap"

	"github.com/coherent-api/contract-poller/shared/go/service_framework/config"
	"github.com/coherent-api/contract-poller/shared/go/service_framework/metrics"
)

type Manager struct {
	// Deal with the manager context
	ctx        context.Context
	cancelFunc context.CancelFunc
	Config     *config.Config

	// Member variables
	logger  *zap.SugaredLogger
	metrics metrics.MetricsInterface
}

type ManagerKey struct{}

func NewManager() (*Manager, error) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	cfg := config.NewConfig()

	var m metrics.MetricsInterface
	var err error
	if cfg.Env == config.Local || cfg.Env == config.Test {
		m, err = metrics.NewNoopMetrics(cfg)
		if err != nil {
			cancelFunc()
			return nil, err
		}
	} else {
		m, err = metrics.NewMetrics(cfg)
		if err != nil {
			cancelFunc()
			return nil, err
		}
	}

	// Setup logger
	var logger *zap.Logger
	switch cfg.Env {
	case config.Test:
		logger = zap.NewNop()
	case config.Local:
		logger, err = zap.NewDevelopment()
		if err != nil {
			cancelFunc()
			return nil, err
		}
	default:
		logger, err = zap.NewProduction()
		if err != nil {
			cancelFunc()
			return nil, err
		}
	}

	manager := &Manager{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		Config:     cfg,
		logger:     logger.Sugar(),
		metrics:    m,
	}

	// Always stat the number of goroutines running in a service
	manager.BackgroundService("stat_num_goroutines", func(ctx context.Context) error {
		for {
			time.Sleep(1 * time.Second)
			manager.Metrics().Gauge("num_goroutines", float64(runtime.NumGoroutine()), []string{}, 1.0)
		}
	})

	return manager, nil
}

func (m *Manager) Context() context.Context {
	return context.WithValue(m.ctx, ManagerKey{}, m)
}

func FromContext(ctx context.Context) *Manager {
	return ctx.Value(ManagerKey{}).(*Manager)
}

func (m *Manager) Metrics() metrics.MetricsInterface {
	return m.metrics
}

func (m *Manager) WaitForInterrupt() {
	notifyChannel := make(chan os.Signal, 1)
	defer close(notifyChannel)
	signal.Notify(notifyChannel, os.Interrupt)
	signal.Notify(notifyChannel, os.Kill)
	select {
	case <-notifyChannel:
		m.cancelFunc()
	case <-m.ctx.Done():
		break
	}
}
