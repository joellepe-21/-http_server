package router

import (
	"fmt"
	"log"
	"server/infrastructure/presentation"
	"server/internal/database"
	"server/internal/domain/repository"
	"server/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (*gin.Engine, error){
	err := database.DbConnect()
	if err != nil{
		log.Fatal(err)
	}

	db := database.DB

	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)

	authController := presentation.NewAuthController(userRepo)
	articleController := presentation.NewArticleController(articleRepo)

	r := gin.Default()

	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"http://localhost:3000"}
    configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
    configCors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	r.Use(middleware.MaxUserMiddleware(100))
	r.Use(cors.New(configCors))

	r.POST("/register", authController.Register)
	r.POST("/authorization", authController.Authorization)

	r.POST("/article", articleController.ArticlePagination)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/add", articleController.Add)
		protected.DELETE("/delete", articleController.Delete)
		protected.PUT("/update", articleController.Update)
	}

	fmt.Println("Маршруты настроены, сервер готов к работе")
	return r, nil
}