package metrics

import (
	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/coherent-api/contract-service/shared/go/service_framework/config"
)

type NoopMetrics struct{}

func NewNoopMetrics(cfg *config.Config) (*NoopMetrics, error) {
	return &NoopMetrics{}, nil
}

// datadog documentation on how to use these functions:
// https://docs.datadoghq.com/metrics/custom_metrics/dogstatsd_metrics_submission/

func (s *NoopMetrics) Incr(name string, tags []string, rate float64) error {
	return nil
}

func (s *NoopMetrics) Decr(name string, tags []string, rate float64) error {
	return nil
}

func (s *NoopMetrics) Count(name string, value int64, tags []string, rate float64) error {
	return nil
}

func (s *NoopMetrics) Gauge(name string, value float64, tags []string, rate float64) error {
	return nil
}

func (s *NoopMetrics) Close() error {
	return nil
}

func (s *NoopMetrics) ServiceCheck(sc *statsd.ServiceCheck) error {
	return nil
}

func (s *NoopMetrics) SimpleEvent(title, text string) error {
	return nil
}

func (s *NoopMetrics) Event(e *statsd.Event) error {
	return nil
}
