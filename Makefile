.PHONY: help build run test test-coverage test-integration test-e2e clean docker-build docker-push migrate-up migrate-down lint swagger

# Variables
APP_NAME=evtaarpro
VERSION?=latest
DOCKER_REGISTRY?=your-registry
DOCKER_IMAGE=$(DOCKER_REGISTRY)/$(APP_NAME):$(VERSION)
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")

help: ## Display this help message
	@echo "EvtaarPro - Makefile Commands"
	@echo "============================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

## Development Commands

run: ## Run the application locally
	@echo "Starting EvtaarPro server..."
	@go run cmd/server/main.go

build: ## Build the application binary
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) cmd/server/main.go
	@echo "Binary created at bin/$(APP_NAME)"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -rf coverage/
	@go clean

## Testing Commands

test: ## Run unit tests
	@echo "Running tests..."
	@go test -v -race -timeout 30s ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@mkdir -p coverage
	@go test -v -race -coverprofile=coverage/coverage.out -covermode=atomic ./...
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report: coverage/coverage.html"

test-integration: ## Run integration tests
	@echo "Running integration tests..."
	@go test -v -race -tags=integration ./testing/...

test-e2e: ## Run E2E tests
	@echo "Running E2E tests..."
	@go test -v -race -tags=e2e ./testing/e2e/...

## Database Commands

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	@go run cmd/migrate/main.go up

migrate-down: ## Rollback database migrations
	@echo "Rolling back migrations..."
	@go run cmd/migrate/main.go down

seed: ## Seed database with test data
	@echo "Seeding database..."
	@go run cmd/seeder/main.go

## Code Quality Commands

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run ./...

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@goimports -w $(GO_FILES)

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

## API Documentation

swagger: ## Generate Swagger documentation
	@echo "Generating Swagger docs..."
	@swag init -g cmd/server/main.go -o api/
	@echo "Swagger docs generated at api/"

## Docker Commands

docker-build: ## Build Docker image
	@echo "Building Docker image: $(DOCKER_IMAGE)"
	@docker build -t $(DOCKER_IMAGE) -f deploy/Dockerfile .

docker-push: ## Push Docker image to registry
	@echo "Pushing Docker image: $(DOCKER_IMAGE)"
	@docker push $(DOCKER_IMAGE)

docker-run: ## Run Docker container locally
	@echo "Running Docker container..."
	@docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

docker-compose-up: ## Start all services with docker-compose
	@echo "Starting services with docker-compose..."
	@docker-compose -f deploy/docker-compose.local.yml up -d

docker-compose-down: ## Stop all services
	@echo "Stopping services..."
	@docker-compose -f deploy/docker-compose.local.yml down

docker-compose-logs: ## View logs from docker-compose
	@docker-compose -f deploy/docker-compose.local.yml logs -f

## Kubernetes Commands

k8s-deploy: ## Deploy to Kubernetes
	@echo "Deploying to Kubernetes..."
	@kubectl apply -f deploy/k8s/

k8s-delete: ## Delete from Kubernetes
	@echo "Deleting from Kubernetes..."
	@kubectl delete -f deploy/k8s/

k8s-logs: ## View pod logs
	@kubectl logs -l app=$(APP_NAME) -f

k8s-restart: ## Restart deployment
	@kubectl rollout restart deployment/$(APP_NAME)

## Monitoring Commands

metrics: ## View application metrics
	@echo "Opening Prometheus..."
	@open http://localhost:9090

grafana: ## Open Grafana dashboard
	@echo "Opening Grafana..."
	@open http://localhost:3000

## Development Tools

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

vendor: ## Vendor dependencies
	@echo "Vendoring dependencies..."
	@go mod vendor

## All-in-One Commands

setup: deps install-tools ## Setup development environment
	@echo "Development environment ready!"

check: fmt vet lint test ## Run all checks

all: clean build test ## Build and test

default: help
