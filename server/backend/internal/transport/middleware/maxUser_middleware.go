package middleware

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func MaxUserMiddleware(maxUser int) gin.HandlerFunc{
	var (
		activeUser int
		mu sync.Mutex
	)
	return func(ctx *gin.Context) {
		mu.Lock()
		log.Println("Active users:", activeUser)
		if activeUser >= maxUser{
			mu.Unlock()
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Server is busy. Please try again later.",
			})
			ctx.Abort()
			return
		}
		activeUser++
		mu.Unlock()

		defer func(){
			mu.Lock()
			activeUser--
			mu.Unlock()
		}()

		ctx.Next()
	}
}