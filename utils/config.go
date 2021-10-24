package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	ListenPort string   `yaml:"port"`
	Database   Database `yaml:"databases"`
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
