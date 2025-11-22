package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"massage": "Missing Authorization Header",
			})
			ctx.Abort()
			return
		}

		if token != "ashish" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"massage": "Invalide Token",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}
