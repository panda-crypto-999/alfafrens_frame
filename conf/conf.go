package conf

import (
	"encoding/json"
	"log"
	"os"
)

type Conf struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

type Mysql struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"dbname"`
	Timeout  string `json:"timeout"`
}

type Redis struct {
	Host      string `json:"host"`
	Password  string `json:"password"`
	Timeout   int    `json:"timeout"`
	MaxActive int    `json:"max_active"`
	MaxIdle   int    `json:"max_idle"`
}

func getConf() Conf {
	buf, err := os.ReadFile("./conf/conf.json")
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	conf := Conf{}
	err = json.Unmarshal(buf, &conf)
	if err != nil {
		log.Panicln("json unmarshal failed: ", err)
	}
	return conf
}

func MysqlConf() Mysql {
	conf := getConf()
	return conf.Mysql
}

func RedisConf() Redis {
	conf := getConf()
	return conf.Redis
}
