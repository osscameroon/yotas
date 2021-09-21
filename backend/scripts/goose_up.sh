#!/bin/bash

ENV_FILE=../.env
if [ $# -ge 1 ]; then
	ENV_FILE=$1
fi

export $(grep -v '#.*' $ENV_FILE | xargs)
PG_STRING="host=$PG_HOST dbname=$PG_DBNAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

cd ../migrations/
goose postgres "$PG_STRING" up
