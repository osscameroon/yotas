.DEFAULT_GOAL := help

## test: run tests over the entire code base
test:
	make test -C ./backend

## lint: run linters over the entire code base
lint:
	make lint -C ./backend

## install-hooks: install hooks
install-hooks:
	ln -s $(PWD)/githooks/pre-push .git/hooks/pre-push

all: help
help: Makefile
	@echo " Choose a command..."
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
