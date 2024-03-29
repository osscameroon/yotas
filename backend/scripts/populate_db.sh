#!/bin/bash

ENV_FILE=../.env
if [ $# -ge 1 ]; then
	ENV_FILE=$1
fi

#install goose
if [[ -z "$(which goose)" ]]; then
	go install github.com/pressly/goose/v3/cmd/goose@latest
fi

#install psql
if [[ -z "$(which psql)" ]]; then
	sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
	wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
	apt-get update
	apt-get -y install postgresql
fi

./goose_up.sh $ENV_FILE

export $(grep -v '#.*' $ENV_FILE | xargs)
PG_STRING="host=$PG_HOST dbname=$PG_DBNAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

cd ../dev/setup-postgres
while read -r file; do
	psql "$PG_STRING" -f "$file"
done < ./sql_file_ordered.txt
