package audit

import (
	"github.com/jinzhu/gorm"
	"connect/conf"
	"fmt"
	"log"
)

var db *gorm.DB

func InitDb(){
	//建立数据库连接
	var db_url string
	db_url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Setting.Audit.DbUser,
		conf.Setting.Audit.DbPassword,
		conf.Setting.Audit.DbHost,
		conf.Setting.Audit.DbPort,
		conf.Setting.Audit.DbName,
			)
	db,err := gorm.Open(conf.Setting.Audit.DbType,db_url)

	if err != nil {
		log.Fatalf("connect to mysql %v err: %v",conf.Setting.Audit.DbType,err)
	}

	//自动初始化表
	db.AutoMigrate(Log{})

	//开连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if conf.Setting.Debug{
		db.LogMode(true)
	}
}