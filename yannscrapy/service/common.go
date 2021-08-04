package service

import (
	"yannscrapy/logger"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	AscSort  = "asc"
	DescSort = "desc"
)

// 文件信息
type FileInfo struct {
	Path       string `json:"path"`
	Size       int64  `json:"size"`
	ModTimeStr string `json:"modTime"`
	modTime    time.Time
}

// 文件比较结构
type FileInfoWrapper struct {
	fileInfo []*FileInfo
	by       func(p, q *FileInfo) bool
}

// 文件内容
type FileContent struct {
	LineNum int64  `json:"linenum"`
	Content string `json:"content"`
}

func (fiw FileInfoWrapper) Len() int {
	return len(fiw.fileInfo)
}

func (fiw FileInfoWrapper) Swap(i, j int) {
	fiw.fileInfo[i], fiw.fileInfo[j] = fiw.fileInfo[j], fiw.fileInfo[i]
}

func (fiw FileInfoWrapper) Less(i, j int) bool {
	return fiw.by(fiw.fileInfo[i], fiw.fileInfo[j])
}

// 转换成正则表达式
func toRegexp(pattern string) string {
	tmp := strings.Split(pattern, ".")
	for i, t := range tmp {
		s := strings.Replace(t, "*", ".*", -1)
		tmp[i] = strings.Replace(s, "?", ".", -1)
	}
	return strings.Join(tmp, "\\.")
}

// 根据文件名过滤
func fileterNameFilter(filePath string, excludes []string) bool {
	fileName := filepath.Base(filePath)
	for _, exclude := range excludes {
		pattern := strings.Trim(exclude, " ")
		if pattern == "" {
			continue
		}
		goPattern := toRegexp(pattern)
		matched, err := regexp.MatchString(goPattern, fileName)
		if matched && err == nil {
			return true
		}
	}
	return false
}

// 字符串转换成数字
func stringToInt64WithDefault(key string, strVal string, defaultVal int64) (int64, error) {
	var err error

	strVal = strings.Trim(strVal, " ")
	if strVal == "" {
		return defaultVal, nil
	}
	val, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		msg := fmt.Sprintf("request param error(%s)", key)
		logger.Errorf(msg)
		return -1, errors.New(msg)
	}
	return val, nil
}

// 字符串转换成数字
func stringToInt64(key string, strVal string) (int64, error) {
	var err error

	strVal = strings.Trim(strVal, " ")
	if strVal == "" {
		msg := fmt.Sprintf("request param error(%s)", key)
		logger.Errorf(msg)
		return -1, errors.New(msg)
	}
	val, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		msg := fmt.Sprintf("request param error(%s)", key)
		logger.Errorf(msg)
		return -1, errors.New(msg)
	}
	return val, nil
}

// 文件排序
func FileTimesort(files []*FileInfo, asc bool) {
	if asc {
		sort.Sort(FileInfoWrapper{files, func(p, q *FileInfo) bool {
			if p.modTime.IsZero() || q.modTime.IsZero() {
				return false
			}
			return p.modTime.Before(q.modTime)
		}})
	} else {
		sort.Sort(FileInfoWrapper{files, func(p, q *FileInfo) bool {
			if p.modTime.IsZero() || q.modTime.IsZero() {
				return false
			}
			return q.modTime.Before(p.modTime)
		}})
	}
}

// 检查匹配字符串
func MatchString(pattern string, ignoreCase bool) func(string) bool {
	pattern = strings.Trim(pattern, " ")
	pattern = toRegexp(pattern)
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	r, err := regexp.Compile(pattern)
	if err != nil {
		logger.Errorf(err.Error())
	}
	return func(str string) bool {
		if err != nil {
			logger.Errorf(err.Error())
			return false
		}
		return r.MatchString(str)
	}
}

func FindAllString(pattern string, ignoreCase bool) func(string) ([]string) {
	pattern = strings.Trim(pattern, " ")
	pattern = toRegexp(pattern)
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	r, err := regexp.Compile(pattern)
	if err != nil {
		logger.Errorf(err.Error())
	}
	return func(str string) ([]string) {
		if err != nil {
			logger.Errorf(err.Error())
			return make([]string, 0)
		}
		return r.FindAllString(str, -1)
	}
}