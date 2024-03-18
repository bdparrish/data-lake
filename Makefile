cwd:=$(shell pwd)
YT:=\033[0;33m
NC:=\033[0m

proto: ## Build protobuf models
	@echo "$(YT)Building protobuf ...$(NC)"
	protoc --go_out=./ ./models/v1/schema.proto

run: proto ## Run main
	@echo "$(YT)Running main.go ...$(NC)"
	CONFIG_FILE=./.env go run main.go