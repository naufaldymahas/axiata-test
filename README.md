# Axiata Golang Test

## Requirement

- [Postgresql](https://www.postgresql.org/download/)
- [Golang](https://golang.org/dl/)

## Usage
Restore dump to your PostgreSQL
```
dump-axiata-202101132323.sql
```

Setting .env before start the program, .env example:
```
DB_HOST=localhost
DB_NAME=axiata
DB_USER=dev
DB_PASSWORD=password
DB_PORT=5432
SECRET_KEY=secret
SIGN_KEY=secret
```
Change Email API First at /src/repository/user_repository.go:55

Run this command on your terminal
```bash
go run main.go
```

## Author
Naufaldy Mahas (naufaldymahas@gmail.com)