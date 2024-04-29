cwd := $(shell pwd)
YT:=\033[0;33m
NC:=\033[0m

.PHONY: docs

all: clean lint test build scan

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Removes all binaries and docker images
	@echo -e "$(YT)Running $(C_NAME) clean ...$(NC)"
	@go clean -cache -testcache
	@rm -rf ./target

docs: ## Generate swagger documentation
	@echo -e "$(YT)Generating $(C_NAME) swagger docs ...$(NC)"
	swag fmt -g ./main.go
	swag init --parseDependency -g ./main.go
	find . -type f -name 'swagger.json' -exec sed -i '' '1s/.*/{"schemes":["http"],"host":"localhost:8080",/g' {} \;

tidy: ## Tidy go mod
	@echo -e "$(YT)Tidying $(C_NAME) ...$(NC)"
	rm -f go.sum
	go mod tidy

mocks: ## Generate all mocks
	@echo -e "$(YT)Generating $(C_NAME) files & mocks ...$(NC)"
	mockery --dir=pkg/repository --name=ObjectRepository --filename=object-repository.go --output=test/mocks/repository --outpkg=repository

lint: tidy docs mocks ## Run linter locally
	@echo -e "${YT}Linting $(C_NAME) files (locally) ...${NC}"
	golangci-lint run -v --timeout 10m

build: lint ## Builds a docker image using a locally built binary
	@echo -e "$(YT)Building $(C_NAME) docker image (using local binary) ...$(NC)"
	env GOOS=linux GOARCH=amd64 go build -o ./target/$(C_NAME) -ldflags=$(LD_FLAGS) main.go
	@docker build --platform linux/amd64 -t $(NAME):$(VERSION) -t $(NAME):latest -f Dockerfile .

test: test-unit ## Run all tests locally

test-unit: lint ## Run unit tests locally
	@echo -e "${YT}Running $(C_NAME) unit tests (locally) ...${NC}"
	@mkdir -p ./target/coverage
	ARRAKIS_CONFIGURATION=$(cwd)/test/configs/test.yaml go test -ldflags=$(LD_FLAGS) -v -count=1 -coverprofile=./target/coverage/coverage.out ./...
	go tool cover -html=./target/coverage/coverage.out -o ./target/coverage/coverage.html

scan: lint ## Run security scans
	@echo "$(YT)Running security scans...$(NC)"
	@gosec ./...
	@go list -json -deps ./... | nancy --skip-update-check -q sleuth
