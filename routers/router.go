package routers

import (
	"net/http"

	"github.com/alvinmdj/mygram-api/database"
	_ "github.com/alvinmdj/mygram-api/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/alvinmdj/mygram-api/handlers"
	"github.com/alvinmdj/mygram-api/middlewares"
	"github.com/alvinmdj/mygram-api/repositories"
	"github.com/alvinmdj/mygram-api/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title           MyGram API
// @version         1.0
// @description     This is service for MyGram.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func StartApp(testDB ...*gorm.DB) *gin.Engine {
	db := database.GetDB(testDB...)

	userRepo := repositories.NewUserRepo(db)
	userSvc := services.NewUserSvc(userRepo)
	userHdl := handlers.NewUserHdl(userSvc)

	socialMediaRepo := repositories.NewSocialMediaRepo(db)
	socialMediaSvc := services.NewSocialMediaSvc(socialMediaRepo)
	socialMediaHdl := handlers.NewSocialMediaHdl(socialMediaSvc)

	photoRepo := repositories.NewPhotoRepo(db)
	photoSvc := services.NewPhotoSvc(photoRepo)
	photoHdl := handlers.NewPhotoHdl(photoSvc)

	commentRepo := repositories.NewCommentRepo(db)
	commentSvc := services.NewCommentSvc(commentRepo)
	commentHdl := handlers.NewCommentHdl(commentSvc)

	r := gin.Default()

	// set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 2 << 20 // 2 MiB

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	v1 := r.Group("/api/v1")
	{
		// user routes
		userRouter := v1.Group("/users")
		{
			userRouter.POST("/register", userHdl.Register)
			userRouter.POST("/login", userHdl.Login)
		}

		// authenticated user only routes
		authenticatedRouter := v1.Group("/")
		{
			authenticatedRouter.Use(middlewares.Authentication())

			// social media routes
			socialMediaRouter := authenticatedRouter.Group("/social-medias")
			{
				socialMediaRouter.GET("", socialMediaHdl.GetAll)
				socialMediaRouter.GET("/:socialMediaId", socialMediaHdl.GetOneById)
				socialMediaRouter.POST("", socialMediaHdl.Create)

				// implement authorization middleware
				socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(testDB...), socialMediaHdl.Update)
				socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(testDB...), socialMediaHdl.Delete)
			}

			// photo routes
			photoRouter := authenticatedRouter.Group("/photos")
			{
				photoRouter.GET("", photoHdl.GetAll)
				photoRouter.GET("/:photoId", photoHdl.GetOneById)

				// implement body size middleware to validate uploaded file size
				photoRouter.POST("", middlewares.BodySizeMiddleware(), photoHdl.Create)

				// implement authorization middleware (+ body size middleware for update handler)
				photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), middlewares.BodySizeMiddleware(), photoHdl.Update)
				photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), photoHdl.Delete)
			}

			commentRouter := authenticatedRouter.Group("/photos/:photoId/comments")
			{
				// implement middleware to find photo by photo id
				commentRouter.Use(middlewares.FindPhoto())

				commentRouter.GET("", commentHdl.GetAll)
				commentRouter.GET("/:commentId", commentHdl.GetOneById)
				commentRouter.POST("", commentHdl.Create)

				// implement authorization middleware
				commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), commentHdl.Update)
				commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), commentHdl.Delete)
			}
		}
	}

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
