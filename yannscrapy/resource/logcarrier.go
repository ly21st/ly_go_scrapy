package resource

import (
	"yannscrapy/service"
	"github.com/gin-gonic/gin"
)

// Logcarrier 查询指定目录下的文件
// @Summary 查询指定目录下的文件
// @Description 返回目录下的文件
// @Tags Logcarrier
// @Accept application/json
// @Produce application/json
// @Param   type     query    string     false        "日志类型"
// @Param   path     query    string     false        "目录路径"
// @Param   excludes      query    string     false        "排除文件，支持通配符"
// @Param   sort    query    string     false        "asc(升序)或desc(降序)"
// @Success 200
// @Router /logcarrier [get]
func SearchFiles(c *gin.Context) {
	service.SearchFiles(c)
}


// Logcarrier 查询实时日志
// @Summary 查询文件内容
// @Description 查询文件内容
// @Tags Logcarrier
// @Accept application/json
// @Produce application/json
// @Param   path     query    string     true        "文件路径"
// @Param   maxline     query    string     true        "查询最大行数"
// @Param   reqline     query    int64     false        "请求行号"
// @Param   direction     query    string     false        "up(上翻)或down(下翻)"
// @Success 200
// @Router /logcarrier/detail [get]
func FileDetail(c *gin.Context) {
	service.FileDetail(c)
}

// Logcarrier 查询实时日志
// @Summary 压缩文件下载
// @Description 下载压缩后的文件
// @Tags Logcarrier
// @Accept application/json
// @Produce application/octet-stream
// @Param   path     query    string     true        "文件路径"
// @Param   compress     query    bool     false        "是否压缩，值为true或false"
// @Param   type     query    string     false        "压缩类型,默认为zip"
// @Success 200
// @Router /logcarrier/download [get]
func FileDownload(c *gin.Context) {
	service.FileDownload(c)
}

// Logcarrier 查询实时日志
// @Summary 搜索某个或多个文件的内容
// @Description 搜索某个或多个文件的内容
// @Tags Logcarrier
// @Accept application/json
// @Produce application/json
// @Param   path     query    string     true        "文件完整路径,多个文件用英文逗号分开"
// @Param   word     query    string     true        "查询字符串"
// @Param   max-match-num  query    int64      false        "最大匹配数"
// @Param   order  query    string      false        "asc或desc，默认desc"
// @Param   first-file  query    string      false        "第一个开始搜索的文件"
// @Param   reqline  query    int64      false        "第一个搜索文件开始查询的行号"
// @Param   ignore-case     query    bool     false        "忽略大小写，值为true或false，默认为true"
// @Success 200
// @Router /logcarrier/search/files [get]
func SearchBaseFile(c *gin.Context) {
	service.SearchBaseFile(c)
}


// Logcarrier 查询实时日志
// @Summary 搜索摸个目录下所有文件的内容
// @Description 搜索某个或多个文件的内容
// @Tags Logcarrier
// @Accept application/json
// @Produce application/json
// @Param   path     query    string     true        "文件完整路径,多个文件用英文逗号分开"
// @Param   word     query    string     true        "查询字符串"
// @Param   max-match-num  query    int64      false        "最大匹配数"
// @Param   order  query    string      false        "asc或desc，默认desc"
// @Param   first-file  query    string      false        "第一个开始搜索的文件"
// @Param   reqline  query    int64      false        "第一个搜索文件开始查询的行号"
// @Param   ignore-case     query    bool     false        "忽略大小写，值为true或false，默认为true"
// @Success 200
// @Router /logcarrier/search/dir [get]
func SearchBaseDir(c *gin.Context) {
	service.SearchBaseDir(c)
}


