package main

import (
	"log"
	"server/internal/app"
	_ "server/internal/transport/controllers"
	"server/config"
	_ "server/docs"
)

// @title           Article API
// @version         1.0
// @description     API for managing articles.
// @host            localhost:8000
// @securityDefinitions.apikey  BearerAuth
// @in              header
// @name            Authorization

func main(){
	conf, err := config.LoadConfig()
	if err != nil{
		log.Fatal("config error:", err)
	}

	app.Run(conf)
}