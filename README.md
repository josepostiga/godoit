# GoDoIt

Simple todo list app made with Go used as a learning project.

## Requirements

- Go 1.21 or higher
- If you want to use the database driver, you need to have a PostgreSQL server running or have Docker installed.

## Installation

1. Clone the repository.
2. Run `go mod download` to download the dependencies.
3. Run `go build` to build the program.
4. Copy the `.env.example` file to `.env` and fill in the values.
5. If using the `postgres` database driver:
   1. If you don't have a running PostgreSQL server, you can run the `docker-compose.yaml` file with `docker compose up -d` to get one, quickly.
   2. Run migrations with a migration tool of your choice (recommended: https://github.com/golang-migrate/migrate).
6. Run the program with `./GoDoIt`.

## Upgrading

1. Pull the latest changes.
2. Run `go mod download` to download the dependencies.
3. Run `go build` to build the program.
4. Ensure that the `.env` file is up-to-date.
5. If using the `postgres` database driver:
   1. Run migrations with a migration tool of your choice (recommended: https://github.com/golang-migrate/migrate).
6. Run the program with `./GoDoIt`.
