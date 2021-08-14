package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"yannscrapy/logger"
)

/*
	复制文件
*/
func CopyFile(dstFileName string, srcFileName string, perm os.FileMode) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		logger.Errorf("open file err = %v", err)
		return
	}

	defer srcFile.Close()

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, perm)
	if err != nil {
		logger.Errorf("open file err = %v", err)
		return
	}

	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func CopyDirSuffixFiles(srcDir string, destDir string, suffix string) error {
	filePaths, err := GetAllSuffixFile(srcDir, suffix)
	if err != nil {
		msg := fmt.Sprintf("getAllYmlFile error:%v", err)
		newErr := errors.New(msg)
		return newErr
	}
	logger.Debugf("filePaths: %v", filePaths)
	// 读取所有input类型配置文件
	for _, fileName := range filePaths {
		filePath := srcDir + string(os.PathSeparator) + fileName
		newFile := destDir + string(os.PathSeparator) + fileName
		err = os.Rename(filePath, newFile)
		if err != nil {
			msg := fmt.Sprintf("Rename error:%v", err)
			logger.Errorf(msg)
		}
	}
	return nil
}

// 读取指定目录下所有指定后缀的文件，返回一个列表，包含所有yml文件
func GetAllSuffixFile(pathName string, suffix string) ([]string, error) {
	var filePaths = make([]string, 0)
	_, err := os.Stat(pathName)
	if err != nil {
		logger.Warnf("", "dir %v is not exist.", pathName)
		return filePaths, nil
	}

	rd, err := ioutil.ReadDir(pathName)
	if err != nil {
		return filePaths, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		}
		fileName := fi.Name()
		if !strings.HasSuffix(fileName, suffix) {
			continue
		}
		filePaths = append(filePaths, fileName)
	}
	return filePaths, nil
}
