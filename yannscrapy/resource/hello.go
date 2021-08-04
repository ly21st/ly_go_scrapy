package resource

import "github.com/gin-gonic/gin"

// Hello 示例接口
// @Summary 打招呼接口
// @Description 返回Hello World
// @Tags Hello
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /hello [get]
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
