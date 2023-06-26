install:
	go get ./...

start:
	go run main.go

test:
	go test -race ./...

cover:
	go test -cover ./...

doc:
	swag init

migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down

local-up:
	docker-compose up -d

local-down:
	docker-compose down