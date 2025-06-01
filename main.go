package main

import (
	"github.com/WesleiDev/api-go-estudo/config"
	"github.com/WesleiDev/api-go-estudo/router"
)

var (
	logger *config.Logger
)

func main(){

	logger = config.GetLogger("main")
	logger.Info("Loading server...")

	//Initialize Configs
	err := config.Init()
	
	if(err != nil){
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize() 
}