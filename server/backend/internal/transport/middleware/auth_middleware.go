package middleware

import (
	"strings"
	"net/http"
	"server/pkg"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == ""{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Missing token"})
			ctx.Abort()
			return
		}
		if strings.HasPrefix(tokenString, "Bearer ") {
            tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        } else {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            ctx.Abort()
            return
        }

		claims, err := pkg.ValidateJWT(tokenString)
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid token"})
			ctx.Abort()
			return	
		}
		ctx.Set("login", claims["login"])
		ctx.Next()
	}
}