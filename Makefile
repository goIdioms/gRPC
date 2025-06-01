.PHONY: proto install-goose migration-create migrate-up migrate-down migrate-status migrate-reset db-create db-drop db-reset

include app.env
export

proto:
	@echo "Generating proto files..."
	@chmod +x scripts/generate-proto.sh
	@./scripts/generate-proto.sh

MIGRATION_DIR=./pkg/database/migrations
DATABASE_URL=postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}:${PGPORT}/${PGDATABASE}?sslmode=disable

migration-create:
	@read -p "Enter migration name: " name; \
	goose -dir ${MIGRATION_DIR} create "$$name" sql

migrate-up:
	goose -dir ${MIGRATION_DIR} postgres "${DATABASE_URL}" up
migrate-down:
	goose -dir ${MIGRATION_DIR} postgres "${DATABASE_URL}" down

migrate-status:
	goose -dir ${MIGRATION_DIR} postgres "${DATABASE_URL}" status

migrate-reset:
	goose -dir ${MIGRATION_DIR} postgres "${DATABASE_URL}" reset

db-create:
	createdb -h ${PGHOST} -p ${PGPORT} -U ${PGUSER} ${PGDATABASE}
db-drop:
	dropdb -h ${PGHOST} -p ${PGPORT} -U ${PGUSER} ${PGDATABASE}
db-reset: db-drop db-create migrate-up