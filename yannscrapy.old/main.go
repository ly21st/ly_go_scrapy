package main

import (
	"flag"
	"fmt"
	"yannscrapy/config"
	"yannscrapy/logger"
	"yannscrapy/resource"
)

var ver = flag.Bool("v", false, "版本号: 1.1.0")

// @title yannscrapy API Docs
// @version 1.0
// @description This is yannscrapy.
// @BasePath /api/v1
func main() {

	flag.Parse()
	if *ver == true {
		fmt.Println("version:1.1.0")
		return
	}

	configFile := config.GetConfigFile()

	// load config
	if err := config.Init(configFile); err != nil {
		panic(err)
	}

	// init logger
	if err := logger.Init(config.Conf.LogConfig); err != nil {
		panic(err)
	}

	resource.InitRouter()

	//web_scraper.Scrapy_main()
}
