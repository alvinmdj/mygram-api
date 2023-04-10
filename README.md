# MyGram REST API

Hacktiv8 Scalable Web Service with Go - Final Project

## Dependencies

- `go get github.com/asaskevich/govalidator`
- `go get github.com/golang-jwt/jwt/v5`
- `go get github.com/gin-gonic/gin`
- `go get golang.org/x/crypto`
- `go get gorm.io/driver/postgres`
- `go get gorm.io/gorm`
- `go get github.com/joho/godotenv`

## Setup DB (Postgres)

- Login psql: `psql -U postgres`
- Show databases: `\list` or `\l`
- Create database: `CREATE DATABASE db_mygram_api;`
- Select database: `\c db_mygram_api`
- Show tables: `dt`

## Run app

`go run .`

## Create random string

`openssl rand -base64 32`
