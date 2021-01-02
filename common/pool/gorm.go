package pool

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	logs "github.com/sirupsen/logrus"
	"qing_novel/config"
	"sync"
	"time"
)

type DefaultPool struct {
	once sync.Once
	Gorm *gorm.DB
}

var _default_content = &DefaultPool{}

func DefaultContent() *DefaultPool {
	_default_content.once.Do(initSqlContent)
	return _default_content
}
func initSqlContent() {
	db, err := NewGorm(config.Config.SqlContent)
	_default_content.Gorm = db
	if err != nil {
		logs.Error("gorm初始化失败")
		return
	}
}

func NewGorm(conf config.SqlConf) (*gorm.DB, error) {
	return newMysql(conf)
}

func newMysql(conf config.SqlConf) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.DataBase)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		logs.Error("mysql连接失败,连接信息:", conn)
		return nil, err
	}

	if conf.Debug {
		//Gorm.LogMode(true)
		//Gorm.SetLogger(logs.Info(os.Stdout, "\r\n", 0))
	}
	db.SingularTable(true)

	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(conf.MaxIdleConns)

	return db, nil
}
