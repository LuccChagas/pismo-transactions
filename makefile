generate-docs:
	swag init -g src/router/handler.go --output docs/app

run-migration:
	go run src/cmd/migration/main.go

start-app:
	go run src/cmd/app/main.go

database-up:
	docker-compose up

database-down:
	docker-compose up