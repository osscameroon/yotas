# Yotas-Backend

This is the backend of the Yotas system.

## Requirements

- golang
- make
- postgres

## How to install/start

- `git clone https://github.com/osscameroon/yotas` to clone the repo
- Copy `.env.example` to `.env` and configure with correct parameters.
- `make docker_db` to pull/run the postgres server with all user/password/database all set.
- `make run` to start the backend server.

- [not required] `make lint_install` to install utils for the linter.
