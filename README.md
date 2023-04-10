# MyGram REST API

Hacktiv8 Scalable Web Service with Go - Final Project

## Endpoints

| Category     | Method | Endpoint                  | Middleware                     | Description               |
|--------------|--------|---------------------------|--------------------------------|---------------------------|
| User         | POST   | /api/v1/users/register    | -                              | User registration         |
| User         | POST   | /api/v1/users/login       | -                              | User login                |
| Photo        | GET    | /api/v1/photos            | Authentication                 | Get all photos            |
| Photo        | GET    | /api/v1/photos/:id        | Authentication                 | Get photo by ID           |
| Photo        | POST   | /api/v1/photos            | Authentication                 | Create new photo          |
| Photo        | PUT    | /api/v1/photos/:id        | Authentication & Authorization | Update photo by ID        |
| Photo        | DELETE | /api/v1/photos/:id        | Authentication & Authorization | Delete photo by ID        |
| Comment      | GET    | /api/v1/comments          | Authentication                 | Get all comments          |
| Comment      | GET    | /api/v1/comments/:id      | Authentication                 | Get comment by ID         |
| Comment      | POST   | /api/v1/comments          | Authentication                 | Create new comment        |
| Comment      | PUT    | /api/v1/comments/:id      | Authentication & Authorization | Update comment by ID      |
| Comment      | DELETE | /api/v1/comments/:id      | Authentication & Authorization | Delete comment by ID      |
| Social Media | GET    | /api/v1/social-medias     | Authentication                 | Get all social medias     |
| Social Media | GET    | /api/v1/social-medias/:id | Authentication                 | Get social media by ID    |
| Social Media | POST   | /api/v1/social-medias     | Authentication                 | Create new social media   |
| Social Media | PUT    | /api/v1/social-medias/:id | Authentication & Authorization | Update social media by ID |
| Social Media | DELETE | /api/v1/social-medias/:id | Authentication & Authorization | Delete social media by ID |

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
