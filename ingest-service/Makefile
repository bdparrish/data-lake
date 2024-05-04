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

scan: test ## Run security scans
	@echo "$(YT)Running security scans...$(NC)"
	@gosec ./...
	@go list -json -deps ./... | nancy --skip-update-check -q sleuth

build: scan ## Builds the docker image
	@echo -e "$(YT)Building ingest-service docker image ...$(NC)"
	@env GOOS=linux GOARCH=amd64 go build -o ./target/ingest-service main.go
	@docker build --platform linux/amd64 -t ingest-service:latest -f Dockerfile .
	@docker scout quickview ingest-service:latest