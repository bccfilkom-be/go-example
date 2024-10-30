include .env

goose_conn=GOOSE_DRIVER=${DB_DRIVER} GOOSE_DBSTRING="postgres://${DB_USERNAME}:${DB_PASSWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}"
migration_dir=./db/migrations/
goose_env=${goose_conn} GOOSE_MIGRATION_DIR=${migration_dir}

run:
	@go run cmd/api/main.go

db/migrate:
	@${goose_env} goose up

db/status:
	@${goose_env} goose status

.PHONY: run db/migrate db/status
