# account-service

## Quick start with docker-compose

```
docker-compose up
```

## Installation

- Install Go 1.18: https://go.dev/doc/install
- Install Postgres: `docker run -d --name=postgres -p 5432:5432 -e POSTGRES_PASSWORD=kidsloop postgres`
- Start Postgres: `docker start postgres`
- Initialize the application and test databases:
    - `docker container exec -it postgres psql -U postgres -c "create database account_service;"`
    - `docker container exec -it postgres psql -U postgres -c "create database account_service_test;"`
    - `POSTGRES_DB=account_service go run scripts/init_test_db.go`
    - `POSTGRES_DB=account_service_test go run scripts/init_test_db.go`
- Copy contents of `.env.example` to `.env`

## Running

- `go run main.go`

## Testing

- `go test ./... -v`

## API Docs

#### Installation

Install the CLI: `go install github.com/swaggo/swag/cmd/swag@latest`

#### Usage

Run the server and browse to http://localhost:8080/swagger/index.html#/

- `swag init` - generate swagger files
- `swag fmt` - format swagger comments
