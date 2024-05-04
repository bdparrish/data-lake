cwd:=$(shell pwd)
YT:=\033[0;33m
NC:=\033[0m

.DEFAULT_GOAL := all

all: test build

tidy: ## Tidy go mod
	@echo -e "$(YT)Tidying $(C_NAME) ...$(NC)"
	rm -f go.sum
	go mod tidy

proto: ## Build protobuf models
	@echo "$(YT)Building protobuf ...$(NC)"
	buf generate

mocks: ## Generate mocks
	@echo "$(YT)Generating mocks ...$(NC)"
	mockery --all --keeptree --output ./test/mocks/ --outpkg mocks

lint: tidy proto mocks ## Run linter locally
	@echo -e "${YT}Linting $(C_NAME) files (locally) ...${NC}"
	@golangci-lint run -v --timeout 10m

test-unit: lint ## Run tests
	@echo "$(YT)Running tests ...$(NC)"
	@CONFIG_FILE=$(cwd)/test/configs/test.yaml go test -v -cover ./pkg/log/... ./pkg/config/...

test-int: lint ## Run integration tests - this will start LocalStack, await healthy LocalStack container, run tests, and clean up docker containers.
	$(eval STACK_NAME:=test-int)
	@echo "$(YT)Starting compose stack '${STACK_NAME}' for integration tests ...$(NC)"
	@docker compose -p ${STACK_NAME} -f ./test/deploy/docker-compose.yml down --remove-orphans -v
	@docker compose -p ${STACK_NAME} -f ./test/deploy/docker-compose.yml up --exit-code-from integration-test-service --remove-orphans
	@docker compose -p ${STACK_NAME} -f ./test/deploy/docker-compose.yml down --remove-orphans -v
	@echo "$(YT)Tests Complete.$(NC)"

test: test-unit test-int ## Run all tests

scan: test ## Run security scans
	@echo "$(YT)Running security scans...$(NC)"
	@gosec ./...
	@go list -json -deps ./... | nancy --skip-update-check -q sleuth

migrate: ## Run database migrations
	@echo -e "$(YT)Running database migrations ...$(NC)"
	@go run main.go migrate

start: ## Start the docker stack
	@echo -e "$(YT)Running data-lake docker stack ...$(NC)"
	@docker compose -p data-lake --env-file deploy/.env -f deploy/docker-compose.yml up --remove-orphans -d

stop: ## Stop the docker stack
	@echo -e "$(YT)Stopping data-lake docker stack ...$(NC)"
	@docker compose -p data-lake -f deploy/docker-compose.yml down --remove-orphans

put-object:
	@echo -e "$(YT)Putting object to S3 ...$(NC)"
	@aws --endpoint-url=http://localhost:4566 s3 cp data/sample.txt s3://data-lake-ingest-bucket/app/data/sample.txt
	@aws --endpoint-url=http://localhost:4566 s3 cp data/sample-2.txt s3://data-lake-ingest-bucket/app/data/sample-2.txt

read-ingest-messages:
	@echo -e "$(YT)Reading ingest messages from SQS ...$(NC)"
	@aws --endpoint-url=http://localhost:4566 sqs receive-message --max-number-of-messages 10 --queue-url http://localhost:4566/000000000000/data-lake-ingest-queue

purge-ingest-messages:
	@echo -e "$(YT)Purging ingest messages from SQS ...$(NC)"
	@aws --endpoint-url=http://localhost:4566 sqs purge-queue --queue-url http://localhost:4566/000000000000/data-lake-ingest-queue

read-log-messages:
	@echo -e "$(YT)Reading log messages from SQS ...$(NC)"
	@aws --endpoint-url=http://localhost:4566 sqs receive-message --max-number-of-messages 10 --queue-url http://localhost:4566/000000000000/data-lake-logger-queue

purge-log-messages:
	@echo -e "$(YT)Purging log messages from SQS ...$(NC)"
	@aws --endpoint-url=http://localhost:4566 sqs purge-queue --queue-url http://localhost:4566/000000000000/data-lake-logger-queue