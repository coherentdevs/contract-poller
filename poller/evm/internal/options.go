package contract_poller

import (
	"golang.org/x/time/rate"
)

type opt func(p *contractPoller)

func WithAbiClient(a AbiClient) opt {
	return func(p *contractPoller) {
		p.abiClient = a
	}
}
func WithRateLimiter(r *rate.Limiter) opt {
	return func(p *contractPoller) {
		p.rateLimiter = r
	}
}
func WithNodeClient(e EvmClient) opt {
	return func(p *contractPoller) {
		p.evmClient = e
	}
}
func WithDatabase(db Database) opt {
	return func(p *contractPoller) {
		p.db = db
	}
}
