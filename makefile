include .env

file_name=
db_path=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

run:
	go run main.go

migrate-create:
	migrate create -ext sql -dir db/migration/ -seq ${file_name}

migrate-up:
	migrate -path db/migration -database ${db_path} -verbose up

migrate-down:
	migrate -path db/migration -database ${db_path} -verbose down

migrate-up-1:
	migrate -path db/migration -database ${db_path} -verbose up 1

migrate-down-1:
	migrate -path db/migration -database ${db_path} -verbose down 1

loadtest-reserve:
	k6 run --out json=output.json --summary-export=summary-export.json loadtest/reserve_ticket.js

.PHONY: migrate-create migrate-up migrate-down migrate-up-1 migrate-down-1 run loadtest-reserve