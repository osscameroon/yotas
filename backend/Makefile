BINARYNAME=yotas
.DEFAULT_GOAL := help

# Database variables
HOSTNAME=localhost
DBNAME=yotas
USER=yotas
PASSWORD=yotas
GOOSE_MIGRATION=cd migrations && goose postgres 'host=$(HOSTNAME) dbname=$(DBNAME) user=$(USER) password=$(PASSWORD) sslmode=disable'


# commands

## --------- application ----------
run:
	go run main.go

## build: build application binary.
build:
	go build -o $(BINARYNAME)

## fmt: format Go source code
fmt:
	go fmt ./...
## --------- application ----------

## --------- migrations ----------
install_goose:
	go get -u github.com/pressly/goose/v3/cmd/goose

goose_up:
	$(GOOSE_MIGRATION) up

goose_down:
	$(GOOSE_MIGRATION) down
## --------- migrations ----------


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
