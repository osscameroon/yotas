-include .env
BINARYNAME=yotas
.DEFAULT_GOAL := help
PG_STRING = 'host=${PG_HOST} dbname=${PG_DBNAME} user=${PG_USER} password=${PG_PASSWORD} sslmode=disable'

# Database variables

# commands
## application
run:
	go run main.go

## build: build application binary.
build:
	go build -o $(BINARYNAME)

## fmt: format Go source code
fmt:
	go fmt ./...

## build-docker: build docker image
build-docker: test
	docker build -t yotas-api .

## migrations
install_goose:
	go get -u github.com/pressly/goose/v3/cmd/goose

goose_up:
	cd migrations && goose postgres $(PG_STRING) up

goose_down:
	cd migrations && goose postgres $(PG_STRING) down

## start-postgres:
start-postgres:
	docker-compose up postgres

## stop-postgres:
stop-postgres:
	docker-compose down

## vet: examines Go source code and reports suspicious constructs
vet:
	go vet ./...

test: vet fmt
	go test ./...

check-lint:
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.8
endif
ifeq (, $(shell which errcheck))
	go get -u github.com/kisielk/errcheck
endif

## lint:
lint: check-lint
	golangci-lint run ./... --timeout 15m0s
	errcheck -exclude ./.golangci-errcheck-exclude.txt ./...

## clean: remove binary releases
clean:
	rm -rf $(BINARYNAME)

all: help
help: Makefile
	@echo " Choose a command..."
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
