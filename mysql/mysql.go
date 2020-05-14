package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" // github.com/go-sql-driver/mysql
	_ "github.com/jmoiron/sqlx"        // github.com/jmoiron/sqlx
)

// db 設定
const (
	userName = "root"
	password = ""
	host     = "127.0.0.1"
	port     = "3306"
	dbName   = "nb_schema"
)

// Initdb : init mysql db function
func Initdb() (*sql.DB, error) {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// "username:password@tcp(host:port)/數據庫?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Printf("%s\n", path)

	// 第一個是 driverName 第二個則是 database 的設定 path
	// 也可以用 var DB *sql.DB
	DB, _ := sql.Open("mysql", path)

	// 設定 database 最大連接數
	DB.SetConnMaxLifetime(100)

	//設定上 database 最大閒置連接數
	DB.SetMaxIdleConns(10)

	// 驗證是否連上 db
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail:", err)
		return nil, err
	}
	fmt.Println("MySQL connnect success!")

	return DB, nil
}