.PHONY: setup init dev migrate-up

POSTGRES_CONTAINER?=indexer_local
POSTGRES_TEST_CONTAINER?=indexer_local_test

setup:
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/vektra/mockery/v2@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/davidbyttow/govips/v2/vips
	cp .env.sample .env

init:
	make setup-githook
	make remove-infras
	docker-compose up -d
	@echo "Waiting for database connection..."
	@while ! docker exec $(POSTGRES_CONTAINER) pg_isready > /dev/null; do \
		sleep 1; \
	done

init_test:
	make remove-infras
	docker-compose up -d postgres postgres_test
	@echo "Waiting for database connection..."
	@while ! docker exec $(POSTGRES_TEST_CONTAINER) pg_isready > /dev/null; do \
		sleep 1; \
	done

seed:
	@docker exec -t indexer_local sh -c "mkdir -p /seed"
	@docker exec -t indexer_local sh -c "rm -rf /seed/*"
	@docker cp migrations/seed indexer_local:/
	@docker exec -t indexer_local sh -c "PGPASSWORD=postgres psql -U postgres -d indexer_local -f /seed/seed.sql"

remove-infras:
	docker-compose down --remove-orphans

api:
	./run-with-arch.sh api

consumer:
	./run-with-arch.sh consumer

indexer:
	./run-with-arch.sh indexer

grpc:
	./run-with-arch.sh grpc

test:
	./test-with-arch.sh

migrate-down:
	sql-migrate down -env=local

gen-swagger:
	swag init -g ./cmd/api/main.go

gen-mock:
	@mockgen -source=./pkg/entity/contract/interface.go -destination=./pkg/entity/contract/mocks/interface.go
	@mockgen -source=./pkg/entity/nft_sale_entity/interface.go -destination=./pkg/entity/nft_sale_entity/mocks/interface.go

setup-githook:
	@echo Setting up softlink pre-commit hooks
	@git config core.hooksPath .githooks/
	@chmod +x .githooks/*
	@echo - done -

abigen:
	cd pkg/grpc/abi_parser/quixotic/gen
	abigen --abi pkg/grpc/abi_parser/quixotic/gen/QuixoticExchange.json --pkg quixotic --out pkg/grpc/abi_parser/quixotic/gen/quixotic.go
