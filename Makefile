# Variables
GO_BIN := go
MAIN_PKG := ./cmd/cli/main.go
BINARY_NAME := ./bin/capital-gains
DOCKER_IMAGE := capital-gains:latest

build: ## Build the application
	@$(GO_BIN) build -o $(BINARY_NAME) $(MAIN_PKG)

test: ## Run tests
	@$(GO_BIN) test ./...

clean: ## Clean up generate files, like binaries and cached tests
	@$(GO_BIN) clean -testcache
	@rm -f $(BINARY_NAME)

docker-build: ## Build Docker image
	@docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run the application inside a Docker container
	@docker run --rm -i $(DOCKER_IMAGE)

## Default target
all: build

help: ## Shows available Makefile commands in a list
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'