local-up:
	@go run cmd/main.go

compose-up:
	docker-compose up -d --build

migration:
	@migrate create -ext sql -dir migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

test:
	@go test -v ./...

update-docs:
	swag init -g cmd/main.go