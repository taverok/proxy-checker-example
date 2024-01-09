run:
	@DOMAIN_CHECK_CONFIG=config/config.yml go run cmd/main.go

lint:
	golangci-lint run

it:
	docker compose up -d
	@DOMAIN_CHECK_CONFIG=config/config.yml go test -count=1 -v ./...
	docker compose down

test:
	go test -count=1 -v -short ./...