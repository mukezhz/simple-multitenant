package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachin-gautam/gin-api/helper"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := helper.ValidateJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		context.Next()
	}
}
