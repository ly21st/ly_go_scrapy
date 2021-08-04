package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// TODO: 待完善

type Config struct {
	Mode       string `yaml:"mode"`
	Port       int    `yaml:"port"`
	*LogConfig `yaml:"logger"`
	SearchPath []SearchPathServiceConfig  `yaml:"search_path"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

type SearchPathServiceConfig struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
	Excludes []string `yaml:"excludes"`
	Sort string `yaml:"sort"`
}

var Conf = new(Config)
var SearchPathConfigMap = make(map[string]SearchPathServiceConfig)

// 初始化配置
func Init(path string) error {

	if f, err := os.Open(path); err != nil {
		return err
	} else {
		yaml.NewDecoder(f).Decode(Conf)
	}

	for _, serviceConfig := range Conf.SearchPath {
		SearchPathConfigMap[serviceConfig.Type] = serviceConfig
	}

	if _, ok := SearchPathConfigMap["default"]; !ok {
		log.Fatal("config [search_path] error")
	}

	return nil
}
