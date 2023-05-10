package chains

import (
	"github.com/coherentopensource/chain-interactor/client/node"
	"github.com/coherentopensource/contract-poller/drivers/optimism"
	"github.com/coherentopensource/go-service-framework/util"
)

func mustInitOptimismDriver(node node.Client, logger util.Logger, cursor uint64, metrics util.Metrics) (*optimism.Driver, error) {
	driver, err := optimism.NewDriver(
		optimism.MustParseConfig(logger),
		node,
		logger,
		metrics,
		cursor,
	)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
