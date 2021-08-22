# Yotas-Backend

This is the backend of the Yotas system.


## Requirements

- golang
- make
- Running on a PostgreSQL Database


## How to install

- Install posgresql for your OS (or use a docker image).
- Set postgresql:
    - Enter the postgresql CLI mode : `sudo -i -u postgres` then hit `psql`
    - create your database : `CREATE DATABASE yotas;`
    - create the default user's credentials for that database : `CREATE USER yotas WITH PASSWORD 'yotas';`
    - add the user to the database : `GRANT ALL ON DATABASE yotas TO yotas;`

- Set up your migrations

## How to start

- `git clone https://github.com/osscameroon/yotas`
- Copy `.env.example` to `.env` and configure with correct parameters
- Make sure the postgreSQL server is up and running
- hit : `make run`
