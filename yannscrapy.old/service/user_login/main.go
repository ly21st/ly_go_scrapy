package main

import (
	"yannscrapy/service/user_login/controller"
	md "yannscrapy/service/user_login/middleware"
	"yannscrapy/service/user_login/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化db
	dbErr := model.InitDBCon()
	if dbErr != nil {
		panic(dbErr)
	}

	defer model.DB.Close()

	// 初始化Gin实例
	router := gin.Default()
	v1 := router.Group("/apis/v1/")
	{
		v1.POST("/register", controller.RegisterUser)
		v1.POST("/login", controller.Login)
	}

	// secure v1
	sv1 := router.Group("/apis/v1/auth/")
	sv1.Use(md.JWTAuth())
	{
		sv1.GET("/time", controller.CheckToken)

	}
	router.Run(":8081")
}
