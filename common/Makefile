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
	@CONFIG_FILE=$(cwd)/test/configs/test.yaml go test -v -cover ./pkg/log/... ./pkg/config/... ./pkg/.

test-int: lint ## Run integration tests - this will start LocalStack, await healthy LocalStack container, run tests, and clean up docker containers.
	$(eval STACK_NAME:=test-int)
	@echo "$(YT)Starting compose stack '${STACK_NAME}' for integration tests ...$(NC)"
	@docker compose -p ${STACK_NAME} -f ./test/deploy/docker-compose.yml down --remove-orphans -v
	@docker compose -p ${STACK_NAME} -f ./test/deploy/docker-compose.yml up --exit-code-from integration-test-service --remove-orphans
	@docker compose -p ${STACK_NAME} -f ./test/deploy/docker-compose.yml down --remove-orphans -v
	@echo "$(YT)Tests Complete.$(NC)"

test: test-unit test-int ## Run all tests

migrate: ## Run database migrations
	@echo -e "$(YT)Running database migrations ...$(NC)"
	@go run main.go migrate

scan: lint ## Run security scans
	@echo "$(YT)Running security scans...$(NC)"
	@gosec ./...
	@go list -json -deps ./... | nancy --skip-update-check -q sleuth