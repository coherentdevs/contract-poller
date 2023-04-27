package chains

import (
	"github.com/coherentopensource/chain-interactor/client/node"
	"github.com/coherentopensource/contract-poller/drivers/ethereum"
	"github.com/coherentopensource/go-service-framework/util"
)

func mustInitEthereumDriver(node node.Client, logger util.Logger, cursor uint64) (*ethereum.Driver, error) {
	driver, err := ethereum.NewDriver(
		ethereum.MustParseConfig(logger),
		node,
		logger,
		cursor,
	)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
