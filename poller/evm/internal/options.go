package contract_poller

import (
	"github.com/datadaodevs/go-service-framework/util"
	"golang.org/x/time/rate"
)

type opt func(p *contractPoller)

func WithABIClient(a ABIClient) opt {
	return func(p *contractPoller) {
		p.abiClient = a
	}
}
func WithRateLimiter(r *rate.Limiter) opt {
	return func(p *contractPoller) {
		p.rateLimiter = r
	}
}
func WithNodeClient(e EVMClient) opt {
	return func(p *contractPoller) {
		p.evmClient = e
	}
}
func WithDatabase(db Database) opt {
	return func(p *contractPoller) {
		p.db = db
	}
}

func WithLogger(l util.Logger) opt {
	return func(p *contractPoller) {
		p.logger = l
	}
}
