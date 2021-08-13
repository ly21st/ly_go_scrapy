package resource

import (
	"yannscrapy/service/user_login/controller"

	"github.com/gin-gonic/gin"
)

// 添加用户
// @Summary 用户注册
// @Description 返回ok
// @Tags Web用户
// @Accept application/json
// @Produce application/json
// @Param   body     body    string     true       "{`userId`: xxx, `password`: xxx}"
// @Success 200
// @Router /web/register [POST]
func WebUserRegister(c *gin.Context) {
	controller.RegisterUser(c)
}

// 用户登录
// Web User Login
// @Summary 用户登录
// @Description 返回ok
// @Tags Web用户
// @Accept application/json
// @Produce application/json
// @Param   body     body    string     true       "{`userId`: xxx, `password`: xxx}"
// @Success 200
// @Router /web/login [POST]
func WebUserLogin(c *gin.Context) {
	controller.Login(c)
}

// 检查token有效性
// @Summary 检查token有效性
// @Description 返回ok
// @Tags Web用户
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /web/check-token [GET]
func CheckToken(c *gin.Context) {
	controller.CheckToken(c)
}
