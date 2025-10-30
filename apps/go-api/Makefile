.PHONY: build run swagger db-up db-down migrate-up migrate-down migrate-status migrate-create

build:
	go build -o ./bin/api main.go

run:
	air

swagger:
	swag init

db-up:
	docker compose up -d

db-down:
	docker compose down

migrate-up:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/homelab?sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/homelab?sslmode=disable" down

migrate-status:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/homelab?sslmode=disable" status

migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir migrations create $$name sql
