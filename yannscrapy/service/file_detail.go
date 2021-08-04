package service

import (
	"yannscrapy/logger"
	"bufio"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

const (
	LineDirectionUp   = "up"
	LineDirectionDown = "down"
)

type FileDetailResponse struct {
	Path        string         `json:"path"`
	MaxLine     int64          `json:"maxline"`
	Direction   string         `json:"direction"`
	ReqLine     int64          `json:"reqline"`
	FileContent []*FileContent `json:"filecontent"`
	ReadLine    int64          `json:"readline"`
}

type FileDetailDes struct {
	Path      string
	MaxLine   int64
	Direction string
	ReqLine   int64
	BeginLine int64
}


// 读取指定行范围的内容，并且在行前面加行号
func (fileDetailDes *FileDetailDes) readLines(fileDetailRsp *FileDetailResponse) (err error) {
	r, err := os.Open(fileDetailDes.Path)
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}
	defer r.Close()

	sc := bufio.NewScanner(r)
	var lineNum int64
	for sc.Scan() {
		lineNum++
		if lineNum < fileDetailDes.BeginLine {
			continue
		}

		lineContent := strconv.FormatInt(lineNum, 10) + " " + sc.Text()
		fileContent := &FileContent{
			LineNum: lineNum,
			Content: lineContent,
		}
		fileDetailRsp.FileContent = append(fileDetailRsp.FileContent, fileContent)
		fileDetailRsp.ReadLine++
		if fileDetailRsp.ReadLine >= fileDetailDes.MaxLine {
			return nil
		}
	}
	return nil
}

// 获取文件内容
func FileDetail(c *gin.Context) {
	var err error
	fileDetailDes, err := requestParamHandle(c)
	if err != nil {
		logger.Errorf(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}
	err = fileDetailDes.beginLineHandle()
	if err != nil {
		logger.Errorf(err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	createResponse := createResponse(fileDetailDes)
	fileDetailDes.readLines(createResponse)

	c.JSON(200, gin.H{
		"data": createResponse,
	})
}

// 请求参数处理
func requestParamHandle(c *gin.Context) (*FileDetailDes, error) {
	var err error

	path := c.Query("path")
	maxLine := c.Query("maxline")
	reqLine := c.Query("reqline")
	direction := c.Query("direction")

	if path == "" || maxLine == "" {
		msg := "request param error"
		logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	fileDetailDes := &FileDetailDes{
		Path:      path,
		Direction: direction,
	}

	fileDetailDes.MaxLine, err = strconv.ParseInt(maxLine, 10, 64)
	if err != nil {
		msg := "request param error(maxline)"
		logger.Errorf(msg)
		return nil, errors.New(msg)
	}

	if reqLine != "" {
		fileDetailDes.ReqLine, err = strconv.ParseInt(reqLine, 10, 64)
		if err != nil {
			msg := "request param error(reqline)"
			logger.Errorf(msg)
			return nil, errors.New(msg)
		}
	} else {
		fileDetailDes.ReqLine = -1
	}

	if direction != "" && (direction != LineDirectionDown || direction != LineDirectionUp) {
		msg := "request param error(direction)"
		logger.Errorf(msg)
		return nil, errors.New(msg)
	}
	if direction == "" {
		fileDetailDes.Direction = LineDirectionDown
	}

	return fileDetailDes, nil
}

// 创建返回对象
func createResponse(fileDetailDes *FileDetailDes) (*FileDetailResponse) {
	fileDetailResponse := &FileDetailResponse{
		Path:        fileDetailDes.Path,
		MaxLine:     fileDetailDes.MaxLine,
		Direction:   fileDetailDes.Direction,
		ReqLine:     fileDetailDes.ReqLine,
		FileContent: make([]*FileContent, 0),
	}
	return fileDetailResponse
}

// 确定上翻或下翻的开始与结束
func (fileDetailDes *FileDetailDes) beginLineHandle() error {
	if fileDetailDes.ReqLine != -1 {
		if fileDetailDes.Direction == LineDirectionDown {
			fileDetailDes.BeginLine = fileDetailDes.ReqLine
		} else {
			fileDetailDes.BeginLine = fileDetailDes.ReqLine - fileDetailDes.MaxLine
			if fileDetailDes.BeginLine <= 1 {
				fileDetailDes.BeginLine = 1
			}
		}
	} else {
		r, err := os.Open(fileDetailDes.Path)
		if err != nil {
			logger.Errorf(err.Error())
			return err
		}
		defer r.Close()

		stat, err := os.Stat(fileDetailDes.Path)
		if err != nil {
			logger.Errorf(err.Error())
			return err
		}
		fileLines, err := ReadLines(r, stat.Size())
		if err != nil && err != io.EOF {
			logger.Errorf(err.Error())
			return err
		}
		fileDetailDes.BeginLine = fileLines - fileDetailDes.MaxLine
		if fileDetailDes.BeginLine <= 1 {
			fileDetailDes.BeginLine = 1
		}
	}

	return nil
}

// 获取文件行数
func ReadLines(r io.Reader, size int64) (lastLine int64, err error) {
	sc := bufio.NewScanner(r)
	var total int64

	for sc.Scan() {
		lastLine++
		total += int64(len(sc.Text()))
		if total >= size {
			break
		}
	}
	return lastLine, io.EOF
}
