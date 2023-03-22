PROJECT_NAME=coherent-api/contract-poller
GO_DIRS=$$(go list -f {{.Dir}} ./... | grep -v ".go")
GO_TARGETS= ./poller/... ./protos/go/... ./shared/...
CURRENT_BRANCH_LATEST_COMMIT = $(shell git rev-parse HEAD)
MAIN_BRANCH_LATEST_COMMIT = $(shell git rev-parse ORIGIN/MAIN)
LINTER_CONFIG_FILE=.golangci.yaml

.PHONY: protos
protos:
	sh scripts/proto.sh

.PHONY: python-format
python-format:
	black ./pipeline

.PHONY: format
format:
	gofmt -e -l -w -s ${GO_DIRS}
	goimports -e -l -w -local github.com/${PROJECT_NAME} ${GO_DIRS}

.PHONY: lint
lint:
	golangci-lint run ${GO_TARGETS} --config ${LINTER_CONFIG_FILE}
	go mod tidy

.PHONY: poller
poller:
	go run poller/evm/cmd/main.go

.PHONY: infra-up
infra-up:
	docker-compose up --build -d

.PHONY: infra-down
infra-down:
	docker-compose down

.PHONY: tests
tests:
	go test poller/evm/internal/contract_poller.go poller/evm/internal/contract_poller_test.go

.PHONY: mocks
mocks:
	mockery --disable-version-string --all --keeptree --case underscore --dir poller/ --output poller/mocks
	mockery --disable-version-string --all --keeptree --case underscore --dir shared/ --output shared/mocks

.PHONY: db-migrate
db-migrate:
	go run poller/scripts/db_migrate/cmd/main.go

.PHONY: fragment-backfiller
fragment-backfiller:
    go run poller/evm/cmd/fragment_backfiller/main.go
