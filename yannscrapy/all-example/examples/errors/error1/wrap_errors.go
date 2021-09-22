package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}

	return buf, nil
}

func ReadConfig() ([]byte, error) {
	//home := os.Getenv("HOME")
	home := "/home"
	config, err := ReadFile(filepath.Join(home, ".setting.xml"))
	return config, errors.WithMessage(err, "would not read config")
}

func main() {
	_, err := ReadConfig()
	if err != nil {
		fmt.Printf("origial error:%T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("hello,world\n")
}
