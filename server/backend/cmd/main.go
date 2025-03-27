package main

import (
	"log"
	"server/internal/app"
	"server/config"
)

func main(){
	conf, err := config.LoadConfig()
	if err != nil{
		log.Fatal("config error:", err)
	}

	app.Run(conf)
}