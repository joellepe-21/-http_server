package cmd

import (
	"log"
	//"net/http"
	//"server/internal/database"
	"server/internal/config"
	"server/infrastructure/router"
)

func Run(){
	conf, err := config.LoadConfig()
	if err != nil{
		log.Fatal("congif error:", err)
	}

	r, err := router.SetupRouter()
	if err != nil{
		log.Fatal(err)
	}

	// if err = database.DbConnect(); err != nil{
	// 	log.Fatal("Ошибка подключния базы данных", err)
	// }
	

	log.Fatal("Ошибка при запуске сервера",r.Run(conf.Port))
}