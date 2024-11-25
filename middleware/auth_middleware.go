package middleware

import (
	"net/http"
	"transaction_app/services"

	"github.com/gin-gonic/gin"
)

func CustomerAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized cookies",
			})
			ctx.Abort()
			return
		}

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
