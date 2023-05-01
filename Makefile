.PHONY: protos
protos:
	sh scripts/proto.sh

.PHONY: run
run:
	go run poller/cmd/main.go

.PHONY: infra-up
infra-up:
	docker-compose up --build -d

.PHONY: infra-down
infra-down:
	docker-compose down
