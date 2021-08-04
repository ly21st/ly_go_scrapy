package service

import (
	"yannscrapy/logger"
	"github.com/gin-gonic/gin"
)

func SearchBaseDir(c *gin.Context) {
	searchContentDes, err := createSearchContentDes(c)
	if err != nil {
		logger.Errorf(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	// 初始化要搜索的文件列表
	err = initSearchFilesForDir(searchContentDes)
	if err != nil {
		logger.Errorf(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	// 文件按时间排序
	FileTimesort(searchContentDes.Files, searchContentDes.Order == AscSort)
	// 创建响应对象
	searchContentResponse := CreateSearchContentResponse()
    // 获取匹配行
	getMatchLines(searchContentDes, searchContentResponse)
	c.JSON(200, gin.H{
		"data": searchContentResponse,
	})
}

// 在目录中初始化要搜索的文件
func initSearchFilesForDir(searchContentDes *SearchContentDes) (error) {
	searchFilesInfo := &SearchFilesInfo{
		SearchFileCommon: SearchFileCommon{
			Type:     "",
			Path:     searchContentDes.Paths,
			Excludes: []string{"*.gz", "*.tar", "*.zip"},
			Sort:     searchContentDes.Order,
			Files:    make([]*FileInfo, 0),
		},
		filter: fileterNameFilter,
	}

	err := searchFilesInfo.searchFilesHandle()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	searchContentDes.Files = searchFilesInfo.Files
	return nil
}
