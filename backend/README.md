# Yotas-Backend

This is the backend of the Yotas system.


## Requirements

- golang, get golang [here](https://golang.org/dl/)
- docker, get docker [here](https://docs.docker.com/get-docker/)
- make
- Running on a PostgreSQL Database


## How to install
Once you have installed docker and the latest version of golang, start the project with:

```terminal
make start-api
```

To teardown the api and its dependencies run:
```terminal
make stop-api
```

#### To run only the postgres database

```terminal
make start-postgres
```

#### To run the api without its dependencies run

```terminal
make run
```

#### To connect to the postgres db run

```terminal
make connect-postgres
```

#### To delete every tables in the db run

```terminal
make clean-postgres
```

#### To stop the postgres database

```terminal
make stop-postgres
```
