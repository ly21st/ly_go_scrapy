package resource

import (
	"fmt"
	"yannscrapy/config"
	_ "yannscrapy/docs"
	"yannscrapy/logger"
	"yannscrapy/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// docs is generated by Swag CLI, you have to import it.

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Credentials", "false")

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Expose-Headers", "*")

		c.Next()
	}
}

// InitRouter init api routers
func InitRouter() {

	if config.Conf.Mode != "" {
		gin.SetMode(config.Conf.Mode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(Cors(), logger.GinLogger(), logger.GinRecovery(true))

	// router.Use(logger.GinRecovery(true))

	initSwagger(router)
	addRouters(router)

	if config.Conf.Port == 0 {
		config.Conf.Port = 8080
	}
	var addr string
	if config.Conf.Address != "" {
		addr = fmt.Sprintf("%v:%v", config.Conf.Address, config.Conf.Port)
	} else {
		addr = fmt.Sprintf(":%v", config.Conf.Port)
	}
	router.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func addRouters(router *gin.Engine) {
	v1Group := router.Group("/api/v1")
	{
		// health group
		healthGroup := v1Group.Group("/health")
		{
			healthGroup.GET("/", Health)
		}

		grabbingTicketGroup := v1Group.Group("/ticket")
		{
			grabbingTicketGroup.GET("/user-list", GetUserList)
			grabbingTicketGroup.POST("/user", AddUser)
			grabbingTicketGroup.DELETE("/user", DeleteUser)
		}

	}

}

func initSwagger(router *gin.Engine) {
	if config.Conf.Swagger {
		var ip string

		if config.Conf.SwagAddress != "" {
			ip = config.Conf.SwagAddress
		} else {
			ip = utils.GetServerIp()
		}
		if config.Conf.Port == 0 {
			config.Conf.Port = 8080
		}
		logger.Infof("swagger ip:%s", ip)
		urlStr := fmt.Sprintf("http://%s:%d/swagger/doc.json", ip, config.Conf.Port)
		url := ginSwagger.URL(urlStr)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
