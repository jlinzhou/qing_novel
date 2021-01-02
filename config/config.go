package config

import (
	"flag"
	logs "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	logSet "qing_novel/common/logs"
)

var configPath  = flag.String("config", "config/config.yaml", "config file abs path")


var Config = PublicConf{}

type PublicConf struct {
	Log        LogConf       `yaml:"log"`
	SqlContent SqlConf       `yaml:"content_sql"`
	Server     ServerConf    `yaml:"server"`

}

type SqlConf struct {
	User         string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	DataBase     string `yaml:"database"`
	Debug        bool   `yaml:"debug"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxLifetime  int    `yaml:"maxLifetime"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	Engine       string `yaml:"engine"`
}
type LogConf struct {
	Filename string `yaml:"filename"`
	Level    string `yaml:"level"`
	Maxsize  int    `yaml:"maxsize"`
	Maxdays  int    `yaml:"maxdays"`
}

type ServerConf struct {
	Mode         string `json:"mode"`
	Port         string `json:"port"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

func init() {
	data, err := ioutil.ReadFile("./config/config.yaml")

	//yaml解析
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		logs.Error(err)
		return
	}
	logSet.New(Config.Log.Filename, Config.Log.Maxsize, 1<<20, Config.Log.Maxdays, false)

}
func Init() {
	data, err := ioutil.ReadFile(*configPath)

	//yaml解析
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		logs.Error(err)
		return
	}
	logSet.New(Config.Log.Filename, Config.Log.Maxsize, 1<<20, Config.Log.Maxdays, false)

}
