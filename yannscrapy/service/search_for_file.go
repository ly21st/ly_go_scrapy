package service

import (
	"yannscrapy/logger"
	"bufio"
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultMaxMatchNum = 100000
)

type FileContentDetail struct {
	FileContent
	Path string `json:"path"`
}

type SearchContentResponse struct {
	MatchLines   []*FileContentDetail `json:"matchlines"`
	LastScanFile string               `json:"lastScanFile"`
	LastScanLine int64                `json:"lastScanLine"`
	ScannedFiles []string             `json:"scannedFiles"`
}

type SearchContentParam struct {
	Paths       string
	Word        string
	MaxMatchNum int64
	ReqLine     int64
	Order       string
	FirstFile   string
	IgnoreCase  bool
}

type SearchContentDes struct {
	SearchContentParam
	Files            []*FileInfo
	CheckMatchString func(str string) bool
	FindAllString func(str string) []string
}

func SearchBaseFile(c *gin.Context) {
	searchContentDes, err := createSearchContentDes(c)
	if err != nil {
		logger.Errorf(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	// 初始化要搜索的文件列表
	initSearchFiles(searchContentDes)
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

// 创建SearchContentDes对象
func createSearchContentDes(c *gin.Context) (*SearchContentDes, error) {
	var err error

	searchContentParam, err := searchContentParamHandle(c)
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}

	searchContentDes := &SearchContentDes{
		SearchContentParam: *searchContentParam,
		Files:              make([]*FileInfo, 0),
		CheckMatchString:   MatchString(searchContentParam.Word, searchContentParam.IgnoreCase),
		FindAllString: FindAllString(searchContentParam.Word, searchContentParam.IgnoreCase),
	}

	return searchContentDes, nil
}

// 初始化要搜索的文件
func initSearchFiles(searchContentDes *SearchContentDes) {
	pathList := strings.Split(searchContentDes.Paths, ",")
	for _, file := range pathList {
		var modTime time.Time
		stat, err := os.Stat(file)
		if err != nil {
			logger.Errorf(err.Error())
			modTime = time.Now()
		} else {
			modTime = stat.ModTime()
		}
		fileInfo := &FileInfo{
			Path:    file,
			modTime: modTime,
		}
		searchContentDes.Files = append(searchContentDes.Files, fileInfo)
	}
}

// 参数处理
func searchContentParamHandle(c *gin.Context) (*SearchContentParam, error) {
	var paths string
	var word string
	var maxMatchNum int64
	var reqLine int64
	var firstFile string
	var order string
	var ignoreCase bool
	var err error

	paths = c.Query("path")
	paths = strings.Trim(paths, " ")
	if paths == "" {
		msg := "request param error(path is empty)"
		logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	word = c.Query("word")
	word = strings.Trim(word, " ")
	if word == "" {
		msg := "request param error(word is empty)"
		logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	maxMatchNumStr := c.Query("max-match-num")
	maxMatchNum, err = stringToInt64WithDefault("max-match-num", maxMatchNumStr, DefaultMaxMatchNum)
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}

	reqLineStr := c.Query("reqline")
	reqLine, err = stringToInt64WithDefault("reqline", reqLineStr, 1)
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}

	firstFile = c.Query("first-file")
	firstFile = strings.Trim(firstFile, " ")

	order = c.Query("order")
	order = strings.Trim(order, " ")
	if order == "" {
		order = DescSort
	}

	if order != DescSort && order != AscSort {
		msg := "request param error(order value error)"
		logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	ignoreCaseStr := c.Query("ignore-case")
	if ignoreCaseStr == "" {
		ignoreCase = true
	} else {
		ignoreCase, err = strconv.ParseBool(ignoreCaseStr)
		if err != nil {
			msg := "request param error(ignore-case)"
			logger.Errorf(msg)
			return nil, errors.New(msg)
		}
	}

	searchContentParam := &SearchContentParam{
		Paths:       paths,
		Word:        word,
		MaxMatchNum: maxMatchNum,
		ReqLine:     reqLine,
		Order:       order,
		FirstFile:   firstFile,
		IgnoreCase:  ignoreCase,
	}
	return searchContentParam, nil
}

// 创建响应对象
func CreateSearchContentResponse() (*SearchContentResponse) {
	searchContentResponse := &SearchContentResponse{
		MatchLines:   make([]*FileContentDetail, 0),
		ScannedFiles: make([]string, 0),
	}
	return searchContentResponse
}

// 查找匹配行,如果指定了firstFile，则从这个指定的文件开始搜素，忽略在这个文件之前的文件内容
func getMatchLines(searchContentDes *SearchContentDes, searchContentResponse *SearchContentResponse) {
	var foundFirstFile bool
	for index, fileInfo := range searchContentDes.Files {
		if int64(len(searchContentResponse.MatchLines)) >= searchContentDes.MaxMatchNum {
			break
		}
		if searchContentDes.FirstFile == "" {
			getMatchLinesInFile(searchContentDes, index, searchContentResponse)
		} else {
			if fileInfo.Path == searchContentDes.FirstFile {
				foundFirstFile = true
			}
			if !foundFirstFile {
				continue
			}
			// 从指定文件指定行开始搜索
			if fileInfo.Path == searchContentDes.FirstFile {
				getMatchLinesInFileSkipLine(searchContentDes, index, searchContentResponse, searchContentDes.ReqLine)
			} else {
				getMatchLinesInFile(searchContentDes, index, searchContentResponse)
			}
		}
	}
}

// 在单个文件中查找匹配行
func getMatchLinesInFile(searchContentDes *SearchContentDes, index int, searchContentResponse *SearchContentResponse) {
	path := searchContentDes.Files[index].Path
	r, err := os.Open(path)
	if err != nil {
		logger.Errorf(err.Error())
		return
	}
	defer r.Close()
	stat, err := os.Stat(path)
	if err != nil {
		logger.Errorf(err.Error())
		return
	}
	fileSize := stat.Size()

	sc := bufio.NewScanner(r)
	var lineNum int64
	var totalSize int64
	for sc.Scan() {
		lineNum++
		text := sc.Text()
		totalSize += int64(len(text))

		matchStrings := searchContentDes.FindAllString(text)
		for _, oldStr := range matchStrings {
			text = strings.ReplaceAll(text, oldStr, "<em>" + oldStr + "</em>")
		}

		if len(matchStrings) != 0 {
			lineContent := strconv.FormatInt(lineNum, 10) + " " + text
			fileContentDetail := &FileContentDetail{
				FileContent: FileContent{
					LineNum: lineNum,
					Content: lineContent,
				},
				Path: path,
			}
			searchContentResponse.MatchLines = append(searchContentResponse.MatchLines, fileContentDetail)
			searchContentResponse.LastScanFile = path
			searchContentResponse.LastScanLine = lineNum
		}

		// 匹配行数达到最大值
		if int64(len(searchContentResponse.MatchLines)) >= searchContentDes.MaxMatchNum {
			break
		}
		// 到达文件结尾
		if totalSize >= fileSize {
			break
		}
	}
	searchContentResponse.ScannedFiles = append(searchContentResponse.ScannedFiles, path)
}

// 在第一个文件中查找匹配行
func getMatchLinesInFileSkipLine(searchContentDes *SearchContentDes, index int, searchContentResponse *SearchContentResponse, skipLine int64) {
	path := searchContentDes.Files[index].Path
	r, err := os.Open(path)
	if err != nil {
		logger.Errorf(err.Error())
		return
	}
	defer r.Close()
	stat, err := os.Stat(path)
	if err != nil {
		logger.Errorf(err.Error())
		return
	}
	fileSize := stat.Size()

	sc := bufio.NewScanner(r)
	var lineNum int64
	var totalSize int64
	for sc.Scan() {
		lineNum++
		text := sc.Text()
		totalSize += int64(len(text))
		// 从第几行开始搜索
		if lineNum < skipLine {
			continue
		}

		matchStrings := searchContentDes.FindAllString(text)
		for _, oldStr := range matchStrings {
			text = strings.ReplaceAll(text, oldStr, "<em>" + oldStr + "</em>")
		}

		if len(matchStrings) != 0 {
			lineContent := strconv.FormatInt(lineNum, 10) + " " + text
			fileContentDetail := &FileContentDetail{
				FileContent: FileContent{
					LineNum: lineNum,
					Content: lineContent,
				},
				Path: path,
			}
			searchContentResponse.MatchLines = append(searchContentResponse.MatchLines, fileContentDetail)
			searchContentResponse.LastScanFile = path
			searchContentResponse.LastScanLine = lineNum
		}

		// 匹配行数达到最大值
		if int64(len(searchContentResponse.MatchLines)) >= searchContentDes.MaxMatchNum {
			break
		}
		// 到达文件结尾
		if totalSize >= fileSize {
			break
		}
	}
	searchContentResponse.ScannedFiles = append(searchContentResponse.ScannedFiles, path)
}
