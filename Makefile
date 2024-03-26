cwd:=$(shell pwd)
YT:=\033[0;33m
NC:=\033[0m

proto: ## Build protobuf models
	@echo "$(YT)Building protobuf ...$(NC)"
	buf generate

mocks: ## Generate mocks
	@echo "$(YT)Generating mocks ...$(NC)"
	mockery --all --keeptree --output ./test/mocks/ --outpkg mocks

run: proto ## Run main
	@echo "$(YT)Running main.go ...$(NC)"
	CONFIG_FILE=./.env go run main.go

test: proto mocks ## Run tests
	@echo "$(YT)Running tests ...$(NC)"
	go test -v -cover ./pkg/config/... ./pkg/ingest/... ./pkg/.

test-int: proto mocks ## Run integration tests - this will start LocalStack, await healthy LocalStack container, run tests, and clean up docker containers.
	$(eval STACK_NAME:=test-int)
	@echo "$(YT)Starting compose stack '${STACK_NAME}' for integration tests ...$(NC)"
	@docker compose -p ${STACK_NAME} -f ./test/deployments/docker-compose.yml down --remove-orphans -v
	@docker compose -p ${STACK_NAME} -f ./test/deployments/docker-compose.yml up --exit-code-from integration-test-service --remove-orphans
	@docker compose -p ${STACK_NAME} -f ./test/deployments/docker-compose.yml down --remove-orphans -v
	@echo "$(YT)Tests Complete.$(NC)"