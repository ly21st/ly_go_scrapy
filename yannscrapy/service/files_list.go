package service

import (
	"yannscrapy/config"
	"yannscrapy/logger"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
	"time"
)


type FileFilterHandle func(filePath string, excludes []string) bool

// 查询文件的返回内容
type SearchFilesResponse struct {
	SearchFileCommon
}

// 每一个查询文件的类型信息
type SearchFilesInfo struct {
	SearchFileCommon

	filter FileFilterHandle
}

type SearchFileCommon struct {
	Type     string      `json:"type"`
	Path     string      `json:"path"`
	Excludes []string    `json:"excludes"`
	Sort     string      `json:"sort"`
	Files    []*FileInfo `json:"files"`
}

// 查询某个目录下的文件
func (searchFilesInfo *SearchFilesInfo) searchFilesHandle() (err error) {
	files, err := ioutil.ReadDir(searchFilesInfo.Path)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if (searchFilesInfo.filter != nil) && searchFilesInfo.filter(file.Name(), searchFilesInfo.Excludes) {
			continue
		}
		fullPath := searchFilesInfo.Path + "/" + file.Name()
		var modTime time.Time
		// 某个文件的错误，可以忽略
		stat, statErr := os.Stat(fullPath)
		if statErr != nil {
			logger.Errorf(statErr.Error())
		} else {
			modTime = stat.ModTime()
		}

		fileInfo := &FileInfo{
			Path:    fullPath,
			Size:    stat.Size(),
			modTime: modTime,
			ModTimeStr: modTime.String(),
		}
		searchFilesInfo.Files = append(searchFilesInfo.Files, fileInfo)
	}
	return
}

// 文件排序
func (searchFilesInfo *SearchFilesInfo) sort() {
	FileTimesort(searchFilesInfo.Files, searchFilesInfo.Sort == AscSort)
}

// 在指定目录查询文件
func SearchFiles(c *gin.Context) {
	searchFilesInfo, err := searchFilesParamHandle(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		logger.Errorf(err.Error())
		return
	}

	err = searchFilesInfo.searchFilesHandle()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	searchFilesInfo.sort()
	searchFilesResponse := &SearchFilesResponse{
		SearchFileCommon: SearchFileCommon{
			Type:     searchFilesInfo.Type,
			Path:     searchFilesInfo.Path,
			Excludes: searchFilesInfo.Excludes,
			Files:    searchFilesInfo.Files,
			Sort:     searchFilesInfo.Sort,
		},
	}
	c.JSON(200, gin.H{
		"data": searchFilesResponse,
	})
}

// 在指定目录查询文件
func searchFilesParamHandle(c *gin.Context) (*SearchFilesInfo, error) {
	serviceType := c.Query("type")
	path := c.Query("path")
	excludesStr := c.Query("excludes")
	sort := c.Query("sort")

	serviceType, path, excludes, sort, err := searchPathConfig(serviceType, path, excludesStr, sort)
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}
	searchFilesInfo := &SearchFilesInfo{
		SearchFileCommon: SearchFileCommon{
			Type:     serviceType,
			Path:     path,
			Excludes: excludes,
			Sort:     sort,
			Files:    make([]*FileInfo, 0),
		},
		filter: fileterNameFilter,
	}

	return searchFilesInfo, nil
}

// 获取配置，如果用户没传递过来，从配置文件读取
func searchPathConfig(service string, path string, excludesStr string, sort string) (string, string, []string, string, error) {
	var excludes []string
	var defaultService = "default"
	var err error

	defaultSearchPathServiceConfig, _ := config.SearchPathConfigMap[defaultService]

	if service == "" {
		service = defaultService
	}

	if path == "" {
		path = config.SearchPathConfigMap[service].Path
	}
	if path == "" {
		path = defaultSearchPathServiceConfig.Path
	}
	if path == "" {
		logger.Error("search_path is empty")
		return "", "", nil, "", err
	}
	if excludesStr == "" {
		excludes = config.SearchPathConfigMap[service].Excludes
		if excludes == nil {
			excludes = defaultSearchPathServiceConfig.Excludes
		}
	} else {
		excludes = strings.Split(excludesStr, ",")
	}
	if excludes == nil {
		excludes = make([]string, 0)
	}

	if sort == "" {
		sort = config.SearchPathConfigMap[service].Sort
	}
	if sort == "" {
		sort = defaultSearchPathServiceConfig.Sort
	}
	if sort == "" {
		sort = DescSort
	}
	if sort != AscSort && sort != DescSort {
		sort = DescSort
	}
	return service, path, excludes, sort, nil
}
