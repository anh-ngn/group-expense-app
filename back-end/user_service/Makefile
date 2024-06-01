# Variables
DB_URL=postgres://anhng:password@localhost:5432/expense_app?sslmode=disable
MIGRATE_CMD=migrate -database $(DB_URL) -path db/migrations

# Targets
.PHONY: all create-migration up down force status run

all: help

help: ## Show this help.
	@echo "Usage: make <target>"
	@echo
	@echo "Targets:"
	@grep '##' $(MAKEFILE_LIST) | grep -v 'grep' | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

create-migration: ## Create a new migration file. Usage: make create-migration name=create_table_users
	@if [ -z "$(name)" ]; then echo "Migration name is required. Usage: make create-migration name=create_table_users"; exit 1; fi
	@migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up: ## Apply all up migrations
	@$(MIGRATE_CMD) up

migrate-down: ## Rollback the last migration
	@$(MIGRATE_CMD) down 1

migrate-force: ## Force the migration to a specific version. Usage: make force version=1
	@if [ -z "$(version)" ]; then echo "Version number is required. Usage: make force version=1"; exit 1; fi
	@$(MIGRATE_CMD) force $(version)

migrate-status: ## Show migration status
	@$(MIGRATE_CMD) version

run: ## Run the server
	@go run main.go

# generate go code from proto files
gen-proto:
	protoc --go_out=. ./proto/*.proto

# generate sqlc code
gen-sqlc:
	sqlc generate
