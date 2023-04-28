package chains

import (
	"github.com/coherentopensource/chain-interactor/client/node"
	"github.com/coherentopensource/go-service-framework/constants"
	poller "github.com/coherentopensource/go-service-framework/contract_poller"
	"github.com/coherentopensource/go-service-framework/util"
)

func MustInitializeDriver(
	chain constants.Blockchain,
	node node.Client,
	logger util.Logger,
	metrics util.Metrics,
	cursor uint64,
) poller.Driver {
	switch chain {
	case constants.Ethereum:
		driver, err := mustInitEthereumDriver(node, logger, cursor, metrics)
		if err != nil {
			logger.Fatalf("could not initialize ethereum driver: %v", err)
		}
		return driver
	}

	logger.Fatal("Unsupported or missing blockchain ID")
	return nil
}
