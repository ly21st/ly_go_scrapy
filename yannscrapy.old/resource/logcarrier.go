package resource

// Filebeat filebeat状态与统计
// @Summary filebeat统计信息
// @Description filebeat统计信息
// @Tags Filebeat
// @Accept application/json
// @Produce application/json
// @Param   ip     query    string     true       "ip地址"
// @Param   port     query    string     true       "端口"
// @Success 200
// @Router /filebeat/stats [get]
// func FilebeatStats(c *gin.Context) {
// 	service.FilebeatStats(c)
// }
