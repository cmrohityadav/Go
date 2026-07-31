# GIN
- https://gin-gonic.com/
# Content

- []()
- []()
- []()
- []()
- []()
- []()
- []()
- []()
- []()


## Installation

```bash
# gin
go get github.com/gin-gonic/gin@latest
# for .env
go get github.com/joho/godotenv@latest

# PostgreSQL (pgx)
go get github.com/jackc/pgx/v5@latest

```

## Db Connection
### Postgres 
```bash
postgres://<USERNAME>:<PASSWORD>@<HOST>:<PORT>/<DATABASE>?sslmode=require
postgres://myuser:mypassword@db.example.com:5432/mydatabase?sslmode=require

- sslmode=disable For Local
- postgres://postgres:123456@localhost:5432/notes_db?sslmode=disable

- sslmode=require For Cloud
```