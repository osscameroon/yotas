BINARYNAME=yotas

.DEFAULT_GOAL := help

lint_install:
	GO111MODULE=off
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

docker_db:
	POSTGRESQL_DATABASE="yotas"
	POSTGRESQL_USERNAME="postgres"
	POSTGRESQL_PASSWORD="postgres"

	docker run --name postgresql_bitnami \
		-e POSTGRESQL_DATABASE=$POSTGRESQL_DATABASE \
		-e POSTGRESQL_PASSWORD=$POSTGRESQL_PASSWORD \
		-e POSTGRESQL_USERNAME=$POSTGRESQL_USERNAME \
		bitnami/postgresql:latest

run:
	go run main.go

## build: build application binary.
build:
	go build -o $(BINARYNAME)

## fmt: format Go source code
fmt:
	go fmt ./...

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
