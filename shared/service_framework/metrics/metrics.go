package metrics

import (
	"fmt"

	"github.com/DataDog/datadog-go/v5/statsd"

	"github.com/coherent-api/contract-poller/shared/service_framework/config"
)

// Metrics wrapper for customized statsd
type Metrics struct {
	Client *statsd.Client
	cfg    *config.Config
}

func NewMetrics(cfg *config.Config) (*Metrics, error) {
	client, err := statsd.New(fmt.Sprintf("%s:%s", cfg.DatadogIP, cfg.DatadogPort))
	if err != nil {
		return nil, err
	}
	return &Metrics{
		cfg:    cfg,
		Client: client,
	}, nil
}

func (s *Metrics) Incr(name string, tags []string, rate float64) error {
	return s.Client.Incr(name, append([]string{fmt.Sprintf("env:%s", string(s.cfg.Env)), fmt.Sprintf("app:%s", s.cfg.AppName)}, tags...), rate)
}

func (s *Metrics) Decr(name string, tags []string, rate float64) error {
	return s.Client.Decr(name, append([]string{fmt.Sprintf("env:%s", string(s.cfg.Env)), fmt.Sprintf("app:%s", s.cfg.AppName)}, tags...), rate)
}

func (s *Metrics) Count(name string, value int64, tags []string, rate float64) error {
	return s.Client.Count(name, value, append([]string{fmt.Sprintf("env:%s", string(s.cfg.Env)), fmt.Sprintf("app:%s", s.cfg.AppName)}, tags...), rate)
}

func (s *Metrics) Gauge(name string, value float64, tags []string, rate float64) error {
	return s.Client.Gauge(name, value, append([]string{fmt.Sprintf("env:%s", string(s.cfg.Env)), fmt.Sprintf("app:%s", s.cfg.AppName)}, tags...), rate)
}

func (s *Metrics) Close() error {
	return s.Close()
}

func (s *Metrics) ServiceCheck(sc *statsd.ServiceCheck) error {
	return s.Client.ServiceCheck(sc)
}

func (s *Metrics) SimpleEvent(title, text string) error {
	return s.Client.SimpleEvent(title, text)
}

func (s *Metrics) Event(e *statsd.Event) error {
	return s.Client.Event(e)
}
