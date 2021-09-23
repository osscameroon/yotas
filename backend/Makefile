.PHONY: build help clean lint test start-postgres goose_up goose_down fmt build-docker
.PHONY: stop-postgres clean-postgres populate-postgres reset-postgres stop-api start-api connect-api
.DEFAULT_GOAL := help

BINARYNAME=yotas

define install_goose
	go get -u github.com/pressly/goose/v3/cmd/goose
endef

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
	$(call install_goose)

goose_up:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./goose_up.sh

goose_down:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./goose_down.sh

## start-api: starts the api and its dependencies in a docker container
start-api:
	docker-compose up

## stop-api: stops the api and its dependencies in a docker container
stop-api:
	docker-compose down && \
	rm -rf postgres-data

## start-postgres: starts postgres sql, setup and populate yotas database
start-postgres:
	docker-compose up start-postgres
	docker-compose logs -f

## connect-postgres: connect to postgres sql
connect-postgres:
	cd ./scripts/ && ./connect_to_postgres.sh

## clean-postgres: delete every tables stored in the yotas database
clean-postgres:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./drop_yotas_tables.sh

## populate-postgres: populate the database with new data
populate-postgres:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./populate_db.sh

## reset-postgres: clean and populate the data base
reset-postgres:
	make clean-postgres && make populate-postgres

## stop-postgres: remove local postgres-data and teardown postgres
stop-postgres:
	docker-compose down && \
	rm -rf postgres-data

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
