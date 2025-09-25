# Makefile for Goose migrations on Windows (MSYS2)

# Variables
MIGRATIONS_FOLDER = db/migrations
DB_URL = "root:Jyoti@123@tcp(127.0.0.1:3306)/auth_dev"

# Create a new migration
# Usage: make migrate-create name=create_session_table
migrate-create:
	@echo "Creating new migration: $(name)"
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

# Apply all pending migrations
# Usage: make migrate-up
migrate-up:
	@echo "Running migrations..."
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) up

# Rollback the last migration
# Usage: make migrate-down
migrate-down:
	@echo "Rolling back last migration..."
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) down
