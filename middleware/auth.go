package middleware

import (
	"context"
	"fmt"
	"net/http"
	"project/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
			ctx.Abort()
			return
		}

		auths := strings.Split(authHeader, " ")
		if len(auths) != 2 {
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
			ctx.Abort()
			return
		}
		var user model.Customer
		data, err := user.DecryptJWT(auths[1])
		fmt.Println(data)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"message": "unauthorized",
			})
			ctx.Abort()
			return
		}
		customerID := int(data["customer_id"].(float64))
		ctxUserID := context.WithValue(ctx.Request.Context(), "customer_id", customerID)
		ctx.Request = ctx.Request.WithContext(ctxUserID)
		ctx.Next()
	}
}
