package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/panda-crypto-999/alfafrens_frame/conf"
)

var db *gorm.DB

func init() {
	mysqlConf := conf.MysqlConf()
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname, mysqlConf.Timeout)
	var err error
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("connect mysql error:" + err.Error())
	}
}
