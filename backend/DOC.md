# Backend Documentation

## Run Requiremnts
* Golang 1.24.2 or higher
* Self created .env file (in backend folder) with fields:
    - APP_PORT
    - DB_HOST
    - DB_PORT
    - DB_DATABASE
    - DB_USERNAME
    - DB_PASSWORD

## Database
* Postgresql

## Development Tools

### Live-Reloading
Air - Live reload for Go apps

Github: https://github.com/air-verse/air

Install command:
```
go install github.com/air-verse/air@latest
```
To use run:

```
air
```
in backend folder.

### Repository Code Generator
sqlc: A SQL Compiler

Github: https://github.com/sqlc-dev/sqlc

Install command:

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
To use run:

```
sqlc generate
```
in backend folder.

### Database Migrations
golang-migrate

Github: https://github.com/golang-migrate/migrate

Download: https://github.com/golang-migrate/migrate/releases

To run use:


```
migrate -path migrations -database "postgres://{username}:{password}@{host}:{port}/{database}?sslmode=disable" {command} {version}
```
in backend folder.
## Folders

- ### cmd
    Entry point for app, main file place.

- ### internal
    Most of application source code.

    - handlers - Handle request, delegate work and return response
    - middlewares - Middlewares functions
    - repositories - Sqlc generated repository pattern to communicate with database
    - server - General purpose like: setting up routes, database connection 
    - services - Business logic and data manipulations

- ### migrations
    Sql schema migrations files. Before adding any new schema updates see https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md (use incrementing integers versioning)

- ### queries
    Sql queries files for generating code using slqc.

- ### initdb
    Tables content initialization queries.

- ### tests
    Tests folder.
