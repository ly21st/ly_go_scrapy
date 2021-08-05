package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// TODO: 待完善

type Config struct {
	Address         string `yaml:"address"`
	Mode            string `yaml:"mode"`
	Port            int    `yaml:"port"`
	DownloadMaxSize int    `yaml:"downloadMaxSize"`
	*LogConfig      `yaml:"logger"`
	Swagger         bool `yaml:"swagger"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

var Conf = new(Config)

// 初始化配置
func Init(path string) error {

	if f, err := os.Open(path); err != nil {
		return err
	} else {
		yaml.NewDecoder(f).Decode(Conf)
	}

	return nil
}

/*
	获取配置文件
*/
func GetConfigFile() string {
	var configFile = "./config.yaml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	return configFile
}
