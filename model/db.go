package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:book@tcp(book_db:3306)/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//マイグレーション
	//db.AutoMigrate(&Book{})


	return db
}

func GetDBConn() *gorm.DB {
	return db
}

// func GetDBConfig() (string, string) {
// 	return config.Env("dialect"), config.Env("datasource")
// }
