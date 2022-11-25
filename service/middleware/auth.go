package middleware

import (
	"net/http"
	"time"

	"github.com/Fachrulmustofa20/go-microservice-user/service/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := utils.VerifyToken(ctx)
		_ = verifyToken
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
				"error":   err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		exp := userData["exp"].(string)
		date, _ := time.Parse(time.RFC3339, exp)
		if date.Before(time.Now()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
				"error":   "token expired. please login again",
			})
			return
		}

		ctx.Next()
	}
}
