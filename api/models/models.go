package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DbInstance struct {
	Dsn string
	Db  *sql.DB
}

func InitDb() *DbInstance {
	var err error

	db := DbInstance{
		Dsn: "root:123456@tcp(192.168.6.120:3306)/demo",
	}
	db.Db, err = sql.Open("mysql", db.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &db
}