# Carrega variáveis do .env
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: server exp db migrate_create_notes_table migrate_up migrate_douwn

server:
	@go run cmd/http/*.go

db:
	@docker compose up -d

migrate_create_notes_table:
	@migrate create -ext sql -dir db/migrations -seq create_notes_table

migrate_create_users_table:
	@migrate create -ext sql -dir db/migrations -seq create_users_table

migrate_up:
	@migrate -database ${DB_CONN_URL} -path db/migrations up

migrate_douwn:
	@migrate -database ${DB_CONN_URL} -path db/migrations down
