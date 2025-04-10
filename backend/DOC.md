# Backend Documentation

## Run Requiremnts
* Golang 1.24.2 or higher
* Self created .env file

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
### Repository Code Generator
sqlc: A SQL Compiler

Github: https://github.com/sqlc-dev/sqlc

Install command:

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### Database Migrations
golang-migrate

Github: https://github.com/golang-migrate/migrate

Download: https://github.com/golang-migrate/migrate/releases

## Folders

### cmd
Entry point for app, main file place.

### internal
Most of application source code.

### migrations
Sql schema migrations files. Before adding any new schema updates see https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md (use incrementing integers versioning)

### queries
Sql queries files.

### tests
Tests folder.
