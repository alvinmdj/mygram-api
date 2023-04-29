# MyGram REST API

Hacktiv8 Scalable Web Service with Go - Final Project

- [Deployed to Railway (might display 'application error' when the free 500 hours is up every month)](https://mygram-api-production.up.railway.app/swagger/index.html)
- [Postman](https://documenter.getpostman.com/view/16534190/2s93XwzjHg)

Techs:

- Gin Framework
- JWT Authentication & Authorization
- PostgreSQL & GORM
- govalidation
- Swagger docs
- Cloudinary file upload

## Swagger

![Swagger docs](https://raw.githubusercontent.com/alvinmdj/mygram-api/main/assets/swagger-screenshot.png "Swagger docs")

## DB Schema

![DB schema](https://raw.githubusercontent.com/alvinmdj/mygram-api/main/assets/db-schema.png "DB schema")

## Endpoints

| Category     | Method | Endpoint                             | Middleware                     | Description               |
|--------------|--------|--------------------------------------|--------------------------------|---------------------------|
| User         | POST   | /api/v1/users/register               | -                              | User registration         |
| User         | POST   | /api/v1/users/login                  | -                              | User login                |
| Photo        | GET    | /api/v1/photos                       | Authentication                 | Get all photos            |
| Photo        | GET    | /api/v1/photos/:id                   | Authentication                 | Get photo by ID           |
| Photo        | POST   | /api/v1/photos                       | Authentication                 | Create new photo          |
| Photo        | PUT    | /api/v1/photos/:id                   | Authentication & Authorization | Update photo by ID        |
| Photo        | DELETE | /api/v1/photos/:id                   | Authentication & Authorization | Delete photo by ID        |
| Comment      | GET    | /api/v1/photos/:photoId/comments     | Authentication                 | Get all comments          |
| Comment      | GET    | /api/v1/photos/:photoId/comments/:id | Authentication                 | Get comment by ID         |
| Comment      | POST   | /api/v1/photos/:photoId/comments     | Authentication                 | Create new comment        |
| Comment      | PUT    | /api/v1/photos/:photoId/comments/:id | Authentication & Authorization | Update comment by ID      |
| Comment      | DELETE | /api/v1/photos/:photoId/comments/:id | Authentication & Authorization | Delete comment by ID      |
| Social Media | GET    | /api/v1/social-medias                | Authentication                 | Get all social medias     |
| Social Media | GET    | /api/v1/social-medias/:id            | Authentication                 | Get social media by ID    |
| Social Media | POST   | /api/v1/social-medias                | Authentication                 | Create new social media   |
| Social Media | PUT    | /api/v1/social-medias/:id            | Authentication & Authorization | Update social media by ID |
| Social Media | DELETE | /api/v1/social-medias/:id            | Authentication & Authorization | Delete social media by ID |

## Links

- [gorm](https://gorm.io/)
- [govalidator](https://github.com/asaskevich/govalidator)
- [swaggo](https://github.com/swaggo/swag)
- [gin upload file](https://gin-gonic.com/docs/examples/upload-file/single-file/)
- [cloudinary go upload file](https://cloudinary.com/documentation/go_image_and_video_upload)
- [cloudinary go example](https://cloudinary.com/documentation/go_integration#complete_sdk_example)
- [cloudinary destroy file](https://cloudinary.com/documentation/image_upload_api_reference#destroy_method)
- [go live reload (air)](https://github.com/cosmtrek/air)
- [example docker setup with go](https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8)

## Dependencies

- `go get github.com/asaskevich/govalidator`
- `go get github.com/golang-jwt/jwt/v5`
- `go get github.com/gin-gonic/gin`
- `go get golang.org/x/crypto`
- `go get gorm.io/driver/postgres`
- `go get gorm.io/gorm`
- `go get github.com/joho/godotenv`
- `go get github.com/swaggo/swag/cmd/swag`
- `go get github.com/swaggo/gin-swagger`
- `go get github.com/swaggo/files`
- `go get github.com/google/uuid`
- `go get github.com/cloudinary/cloudinary-go/v2`
- `go get github.com/cloudinary/cloudinary-go/v2/api/uploader`

## Setup DB (Postgres)

- Login psql: `psql -U postgres`
- Show databases: `\list` or `\l`
- Create database: `CREATE DATABASE db_mygram_api;`
- Select database: `\c db_mygram_api`
- Show tables: `dt`

## Setup env

- Copy `cp .env.example .env`
- Setup environment variables

## Init swagger docs

`swag init -g routers/router.go`

## Run app

- local development: `go run .` or `air` for live reload
- or using Docker: `docker compose up`

## Create random string

`openssl rand -base64 32`
