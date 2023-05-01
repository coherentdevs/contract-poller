# contract-poller

Contract Poller is a service to fetch all smart contracts on an EVM chain.
It is responsible for determining contract deployment transactions by polling transactions,
receipts and traces using a RPC node, and fetches its relevant ABI from etherscan
and contract metadata (standard, symbol, name, decimals, etc) via a RPC node.

## Installation

To install contract-poller, make sure you have Go installed on your machine. Then, run:

`go get github.com/coherentopensource/contract-poller`


## Usage
Before running contract poller, make sure to have the following values as your environment variables.
```dotenv
ENV=local
BLOCKCHAIN=ethereum

REDIS_HOST=localhost:6379
NODE_HOST={rpc_node_host}

DB_HOST={db_host}

DB_USER={db_user}
DB_PASSWORD={db_password}
DB_NAME={db_name}
DB_PORT=5432
SSL_MODE=disable

ETHERSCAN_API_KEY={etherscan_api_key}

POLLER_AUTO_START=true
# BATCH_SIZE is how many blocks you poll at once
BATCH_SIZE=
# POLLER_POOL_BANDWIDTH should be 3 * BATCH_SIZE
FETCHER_POOL_BANDWIDTH=
# ACCUMULATOR_POOL_BANDWIDTH should be equal to BATCH_SIZE
ACCUMULATOR_POOL_BANDWIDTH=
# WRITER_POOL_BANDWIDTH should be 3 * BATCH_SIZE
WRITER_POOL_BANDWIDTH=
```
To run {Service Name}, simply run (if run locally, it uses .env file):

`make run`

## Contributing

If you would like to contribute to Contract Poller, please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b username/{your-feature-name}`)
3. Make your changes
4. Commit your changes (`git commit -m "Add some feature"`)
5. Push to the branch (`git push origin username/{your-feature-name}`)
6. Create a new Pull Request

## License

Contract Poller is released under the MIT license. See [LICENSE](LICENSE) for more information.
