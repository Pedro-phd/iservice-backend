#variables
include .env
# DATA_BASE_URL=

# Tasks
default: run

run:
	@cd cmd/iservice && go run main.go

docker-up:
	@sudo docker compose up

docker-down:
	@sudo docker compose down

migrate-up:
	@migrate -database $(DB_URL_MIGRATE) -path src/infra/migrations -verbose up

migrate-down:
	@migrate -database $(DB_URL_MIGRATE) -path src/infra/migrations -verbose down
