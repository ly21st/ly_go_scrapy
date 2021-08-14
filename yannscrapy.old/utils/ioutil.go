package utils

import (
	"io"
	"io/ioutil"
	"os/exec"
	"yannscrapy/logger"
)

func RunCmdAndGetResult(cmdStr string) (string, error) {
	var stdout io.ReadCloser
	var err error

	cmd := exec.Command("bash", "-c", cmdStr)
	if stdout, err = cmd.StdoutPipe(); err != nil {
		logger.Error(err)
		return "", err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		logger.Error(err)
		return "", err
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		logger.Error(err)
		return "", err
	} else {
		return string(opBytes), nil
	}
}
