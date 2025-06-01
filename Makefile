.PHONY: proto install-goose migration-create migrate-up migrate-down migrate-status migrate-reset db-create db-drop db-reset

proto:
	@echo "Generating proto files..."
	@chmod +x scripts/generate-proto.sh
	@./scripts/generate-proto.sh

MIGRATION_DIR=./migrations
DB_DSN=postgres://postgres:postgres@localhost:5432/user_service?sslmode=disable

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migration-create:
	@read -p "Enter migration name: " name; \
	goose -dir ${MIGRATION_DIR} create "$$name" sql

migrate-up:
	goose -dir ${MIGRATION_DIR} postgres "${DB_DSN}" up
migrate-down:
	goose -dir ${MIGRATION_DIR} postgres "${DB_DSN}" down

migrate-status:
	goose -dir ${MIGRATION_DIR} postgres "${DB_DSN}" status

migrate-reset:
	goose -dir ${MIGRATION_DIR} postgres "${DB_DSN}" reset

db-create:
	createdb -h localhost -p 5432 -U postgres user_service
db-drop:
	dropdb -h localhost -p 5432 -U postgres user_service
db-reset: db-drop db-create migrate-up