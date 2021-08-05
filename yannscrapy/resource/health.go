package resource

import (
	"yannscrapy/logger"

	"github.com/gin-gonic/gin"
)

// Health 健康接口
// @Summary 健康接口
// @Description 返回ok
// @Tags Health
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /health [get]
func Health(c *gin.Context) {
	logger.Info("hello world")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
