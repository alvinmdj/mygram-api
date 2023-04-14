package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BodySizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		maxBodyBytes := int64(2 << 20) // 2 MiB

		var w http.ResponseWriter = c.Writer
		c.Request.Body = http.MaxBytesReader(w, c.Request.Body, maxBodyBytes)

		if c.Request.ContentLength > maxBodyBytes {
			c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{
				"error":   "REQUEST ENTITY TOO LARGE",
				"message": "request body (file uploaded) too large",
			})
			return
		}

		c.Next()
	}
}
