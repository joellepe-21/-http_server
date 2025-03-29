package app

import (
	"fmt"
	"log"
	"server/config"
	"server/internal/database"
	"server/internal/domain/repository"
	"server/internal/transport/controllers"
	"server/internal/transport/middleware"
	"server/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func Run(confing *config.Config){
	err := database.DbConnect()
	if err != nil{
		log.Fatal(err)
	}

	db := database.DB

	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)

	useCase := usecase.NewArticleUsecase(*articleRepo)

	authController := controllers.NewAuthController(userRepo)
	articleController := controllers.NewArticleController(articleRepo, useCase)

	r := gin.Default()

	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"http://localhost:3000"}
    configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
    configCors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	r.Use(middleware.MaxUserMiddleware(100))
	r.Use(cors.New(configCors))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	// if err = database.DbConnect(); err != nil{
	// 	log.Fatal("Ошибка подключния базы данных", err)
	// }
	
	log.Fatal("Ошибка при запуске сервера",r.Run(confing.Port))
	
}