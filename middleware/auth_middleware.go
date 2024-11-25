package middleware

import (
	"net/http"
	"strings"
	"transaction_app/services"

	"github.com/gin-gonic/gin"
)

func CustomerAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization format",
			})
			ctx.Abort()
			return
		}

		token := tokenParts[1]
		claims, err := services.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}

		authCustomerID, ok := claims["customer_id"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("customer_id", int(authCustomerID))

		ctx.Next()
	}
}
