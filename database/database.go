package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// database - setting
const (
	userName = "root"
	password = ""
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "wristband"
)

// Initdb : init mysql db function
func Initdb() (*gorm.DB, error) {
	openURL := userName + ":" + password + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", openURL)

	if err != nil {
		fmt.Println("opon database fail:", err)
		return nil, err
	}

	// 設定 database 最大連接數
	db.DB().SetMaxOpenConns(100)

	// 設定上 database 最大閒置連接數
	db.DB().SetMaxIdleConns(10)

	// 驗證是否連上 database
	if err := db.DB().Ping(); err != nil {
		fmt.Println("opon database fail:", err)
		return nil, err
	}
	// fmt.Println("database connnect success!")

	return db, nil
}
