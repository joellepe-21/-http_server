package config

import (
	"encoding/json"
	"os"
)

type Config struct{
	Port string `json:"port"`
}

func LoadConfig() (*Config, error){
	file, err := os.ReadFile("internal/config/config.json")
	if err != nil{
		return nil, err
	}
	
	var conf Config

	err = json.Unmarshal(file, &conf)
	if err != nil{
		return nil, err
	}

	if conf.Port == ""{
		conf.Port = ":8080"
	}

	return &conf, nil
}