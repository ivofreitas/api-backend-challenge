install:
	go get ./...

start:
	go run main.go

test:
	go test -race ./...

cover:
	go test -cover ./...

local-up:
	docker-compose up -d

local-down:
	docker-compose down --remove-orphans

configure:
	cp config/example.env config/.env
	swag init