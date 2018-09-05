package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/makki0205/config"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open("mysql", config.Env("DB_USER")+":"+config.Env("DB_PASS")+"@tcp("+config.Env("DB_CONN")+")/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})


	return db
}

func GetDBConn() *gorm.DB {
	return db
}

// func GetDBConfig() (string, string) {
// 	return config.Env("dialect"), config.Env("datasource")
// }
