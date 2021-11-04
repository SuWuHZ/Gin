package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Server   ServerInfo `yaml:"server"`
	Database Database   `yaml:"database"`
}

type ServerInfo struct {
	Port int `yaml:"port"`
}

type Database struct {
	SqlConn  string `yaml:"sqlconn"`
	MongoUri string `yaml:"mongouri"`
}

var ServerConf ServerConfig

func init() {

	yamlFile, err := ioutil.ReadFile("cfg/config.yaml")
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &ServerConf)
	if err != nil {
		panic(err.Error())
	}
}
