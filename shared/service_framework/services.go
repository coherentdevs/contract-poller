package service_framework

import (
	"context"
	"time"
)

func (m *Manager) PeriodicService(name string, job func(ctx context.Context) error, duration time.Duration) {
	m.Logger().Infof("starting %s", name)
	go func() {
		ticker := time.NewTicker(duration)

		for {
			select {
			case <-ticker.C:
				if err := job(m.ctx); err != nil {
					m.Logger().Errorf("error in background job %s: %v", name, err)
				}
			case <-m.ctx.Done():
				break
			}
		}
	}()
}

func (m *Manager) BackgroundService(name string, job func(ctx context.Context) error) {
	m.Logger().Infof("starting %s", name)
	go func() {
		for {
			if err := job(m.ctx); err != nil {
				m.Logger().Errorf("error in background job %s: %v", name, err)
			}
		}
	}()
}
