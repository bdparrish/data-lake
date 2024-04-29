cwd:=$(shell pwd)
YT:=\033[0;33m
NC:=\033[0m

.DEFAULT_GOAL := all

all: test build

tidy: ## Tidy go mod
	@echo -e "$(YT)Tidying $(C_NAME) ...$(NC)"
	rm -f go.sum
	go mod tidy

mocks: ## Generate mocks
	@echo "$(YT)Generating mocks ...$(NC)"
	mockery --all --keeptree --output ./test/mocks/ --outpkg mocks

lint: tidy mocks ## Run linter locally
	@echo -e "${YT}Linting $(C_NAME) files (locally) ...${NC}"
	@golangci-lint run -v --timeout 10m

run: tidy mocks ## Run main
	@echo "$(YT)Running main.go ...$(NC)"
	@CONFIG_FILE=./.env go run main.go

test-unit: lint ## Run tests
	@echo "$(YT)Running tests ...$(NC)"
	@CONFIG_FILE=$(cwd)/test/configs/test.yaml go test -v -cover ./pkg/...

test: test-unit ## Run all tests

scan: lint ## Run security scans
	@echo "$(YT)Running security scans...$(NC)"
	@gosec ./...
	@go list -json -deps ./... | nancy --skip-update-check -q sleuth

build: scan ## Builds the docker image
	@echo -e "$(YT)Building data-lake docker image ...$(NC)"
	@env GOOS=linux GOARCH=amd64 go build -o ./target/data-lake main.go
	@docker build --platform linux/amd64 -t data-lake:latest -f deploy/Dockerfile .
	@docker scout quickview data-lake:latest

start: ## Start the docker stack
	@echo -e "$(YT)Running data-lake docker stack ...$(NC)"
	@docker compose -p data-lake --env-file deploy/.env -f deploy/docker-compose.yml up --remove-orphans

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