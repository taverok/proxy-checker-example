run:
	@ DOMAIN_CHECK_CONFIG=config/config.yml go run cmd/web/main.go

lint:
	@ golangci-lint run

build-linux:
	@ GOOS=linux GOARCH=amd64 go build -o ./bin/checker ./cmd/cli/main.go

it:
	@ docker compose up -d
	@ DOMAIN_CHECK_CONFIG=config/config.yml go test -count=1 -v ./...
	@ docker compose down

test:
	@ go test -count=1 -v -short ./...