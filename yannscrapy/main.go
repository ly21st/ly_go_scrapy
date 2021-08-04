package main

import (
	"yannscrapy/config"
	"yannscrapy/logger"
	"yannscrapy/resource"
	"os"
)

// @title SuperAgent API Docs
// @version 1.0
// @description This is bigdata SuperAgent.
// @BasePath /api/v1
func main() {
	var configFile = "./config.yaml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	// load config
	if err := config.Init(configFile); err != nil {
		panic(err)
	}

	// init logger
	if err := logger.Init(config.Conf.LogConfig); err != nil {
		panic(err)
	}

	resource.InitRouter()

}
