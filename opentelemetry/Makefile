include .env

goose_conn=GOOSE_DRIVER=${DB_DRIVER} GOOSE_DBSTRING="postgres://${DB_USERNAME}:${DB_PASSWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}"
migration_dir=./db/migrations/
goose_env=${goose_conn} GOOSE_MIGRATION_DIR=${migration_dir}

run:
	@go run cmd/api/main.go

seed:
	@go run cmd/seeder/main.go

goose/up:
	@${goose_env} goose up

goose/status:
	@${goose_env} goose status

sqlc/gen:
	@cd db && sqlc generate

test:
	@cd test && hurl --test --very-verbose http.hurl

compose/up:
	@docker compose up -d

compose/down:
	@docker compose down

compose/fresh:
	@docker compose down
	@docker compose up -d --build
	# Waiting for database uptime
	@sleep 4
	@${goose_env} goose up
	# Seeding database
	@go run cmd/seeder/main.go

compose/restart:
	@docker compose down
	@docker compose up -d
	# Waiting for database uptime
	@sleep 4
	@${goose_env} goose up
	# Seeding database
	@go run cmd/seeder/main.go

k6/run:
	@k6 run test/script.js

.PHONY: run db/migrate db/status sqlc/gen test seed compose/up compose/down compose/fresh compose/restart k6/run
