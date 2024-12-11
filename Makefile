local-up:
	@go run cmd/main.go

db-up:
	docker-compose up -d --build

db-down:
	docker-compose down

docker-up:
	docker-compose -f docker-compose-app.yaml up -d --build

docker-down:
	docker-compose -f docker-compose-app.yaml down

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