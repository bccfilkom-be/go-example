include .env

goose_conn=GOOSE_DRIVER=${DB_DRIVER} GOOSE_DBSTRING="postgres://${DB_USERNAME}:${DB_PASSWD}@${DB_HOST}:${DB_PORT}/${DB_NAME}"
migration_dir=./db/migrations/
goose_env=${goose_conn} GOOSE_MIGRATION_DIR=${migration_dir}

sqlc/gen:
	@cd db && sqlc generate

goose/gen:
	@cd db/migrations && goose create ${x} sql

goose/up:
	@${goose_env} goose up

goose/status:
	@${goose_env} goose status

postgres:
	@docker run \
		--name postgres \
		--publish 5432:5432 \
		--env POSTGRES_USER=${DB_NAME} \
		--env POSTGRES_PASSWORD=${DB_PASSWD} \
		--detach postgres:17.0-alpine3.20

.PHONY:
	sqlc/gen
	goose/gen
	goose/up
	goose/status
