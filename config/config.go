package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// 总配置项
type config struct {
	Server       server       `yaml:"server"`
	Database     database     `yaml:"database"`
	Redis        redis        `yaml:"redis"`
	FileSettings fileSettings `yaml:"fileSettings"`
	Log          log          `yaml:"log"`
}

// WEB服务配置项
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置项
type database struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// Redis配置项
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// 文件上传配置项
type fileSettings struct {
	UploadDir string `yaml:"uploadDir"`
	Host      string `yaml:"host"`
}

// 日志配置项
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

var Config *config

// 初始化全局配置项
func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
