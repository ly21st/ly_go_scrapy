package resource

import (
	"github.com/gin-gonic/gin"
	"yannscrapy/service/grabbing_ticket_service"
)

// User management
// @Summary 用户列表
// @Description 返回ok
// @Tags User management
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /ticket/user-list [get]
func GetUserList(c *gin.Context) {
	grabbing_ticket_service.GetUserList(c)
}

// 添加用户
// User management
// @Summary 添加用户
// @Description 返回ok
// @Tags User management
// @Accept application/json
// @Produce application/json
// @Param   body     body    string     true       "{`userId`: xxx, `password`: xxx}"
// @Success 200
// @Router /ticket/user [POST]
func AddUser(c *gin.Context) {
	grabbing_ticket_service.AddUser(c)
}

// 删除用户
// User management
// @Summary 删除用户
// @Description 返回ok
// @Tags User management
// @Accept application/json
// @Produce application/json
// @Param   body     body    string     true       "{`userId`: xxx}"
// @Success 200
// @Router /ticket/user [delete]
func DeleteUser(c *gin.Context) {
	grabbing_ticket_service.DeleteUser(c)
}