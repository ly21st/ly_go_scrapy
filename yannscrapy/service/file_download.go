package service

import (
	"yannscrapy/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const (
	DefaultCompressType = "zip"
	DefaultContentType  = "application/octet-stream"
	ContentTypeZip      = "application/zip"
)

type CompressFunc func(srcFile string, destZip string) error

type FileDownloadDes struct {
	path           string
	compressType   string
	compress       bool
	compressMethod map[string]CompressFunc
	contentType    map[string]string
}

var defaultCompressMethod = Zip
var mutex = &sync.Mutex{}
var glbFileSeq int64 = 0

// 实现下载文件函数
func (fileDownloadDes *FileDownloadDes) downloadFile(c *gin.Context) {
	var filename string
	var destFile string
	if fileDownloadDes.compress {
		filename = filepath.Base(fileDownloadDes.path) + "." + fileDownloadDes.compressType
	} else {
		filename = filepath.Base(fileDownloadDes.path)
	}
	fileContentDisposition := "attachment; filename=" + "\"" + filename + "\""
	c.Writer.Header().Add("Content-Disposition", fileContentDisposition)
	c.Writer.Header().Add("Content-Type", getContentType(fileDownloadDes))
	if ! fileDownloadDes.compress {
		c.File(fileDownloadDes.path)
	} else {
		compressFunc := getCompressMethod(fileDownloadDes)
		destFile = generateCompressName(fileDownloadDes)
		err := compressFunc(fileDownloadDes.path, destFile)
		if err != nil {
			logger.Errorf(err.Error())
			c.File(fileDownloadDes.path)
			return
		}
		defer os.Remove(destFile)
		c.File(destFile)
	}
}

// 下载文件
func FileDownload(c *gin.Context) {
	var path string
	var compress bool
	var compressType string
	var err error
	path = c.Query("path")
	if path == "" {
		msg := "request param error(missing path)"
		logger.Errorf(msg)
		c.JSON(400, gin.H{
			"message": msg,
		})
		return
	}
	compressStr := c.Query("compress")
	if compressStr == "" {
		compress = false
	} else {
		compress, err = strconv.ParseBool(compressStr)
		if err != nil {
			msg := "request param error(compress)"
			logger.Errorf(msg)
			c.JSON(400, gin.H{
				"message": msg,
			})
			return
		}
	}
	compressType = c.Query("type")
	if compressType == "" {
		compressType = DefaultCompressType
	}

	fileDownloadDes := &FileDownloadDes{
		path:           path,
		compress:       compress,
		compressType:   compressType,
		compressMethod: make(map[string]CompressFunc),
		contentType:    make(map[string]string),
	}
	fileDownloadDes.compressMethod[DefaultCompressType] = defaultCompressMethod
	fileDownloadDes.contentType[DefaultCompressType] = ContentTypeZip
	fileDownloadDes.downloadFile(c)
}

// 产生临时压缩文件名
func generateCompressName(fileDownloadDes *FileDownloadDes) string {
	mutex.Lock()
	glbFileSeq += 1
	mutex.Unlock()
	filename := fmt.Sprintf("%s.%d.%d.%s", filepath.Base(fileDownloadDes.path), time.Now().UnixNano(), glbFileSeq, fileDownloadDes.compressType)
	return filename
}

// 获取压缩方法
func getCompressMethod(fileDownloadDes *FileDownloadDes) CompressFunc {
	compressFunc, ok := fileDownloadDes.compressMethod[fileDownloadDes.compressType]
	if !ok {
		compressFunc = defaultCompressMethod
	}
	return compressFunc
}

func getContentType(fileDownloadDes *FileDownloadDes) string {
	contentType, ok := fileDownloadDes.contentType[fileDownloadDes.compressType]
	if !ok {
		contentType = DefaultContentType
	}
	return contentType
}
