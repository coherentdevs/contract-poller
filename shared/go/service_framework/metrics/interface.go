package metrics

import "github.com/DataDog/datadog-go/v5/statsd"

type MetricsInterface interface {
	Incr(name string, tags []string, rate float64) error
	Decr(name string, tags []string, rate float64) error
	Count(name string, value int64, tags []string, rate float64) error
	Gauge(name string, value float64, tags []string, rate float64) error
	Close() error
	ServiceCheck(sc *statsd.ServiceCheck) error
	SimpleEvent(title, text string) error
	Event(e *statsd.Event) error
}
