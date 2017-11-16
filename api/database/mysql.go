package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		log.Fatal(err.Error())
	}
}
