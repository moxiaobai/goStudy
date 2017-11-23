package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/moxiaobai/goStudy/api/config"
)

var SqlDB *sql.DB

func init() {
	config := config.GetConfig()

	var err error
	SqlDB, err = sql.Open(config.Database.Driver, config.Database.Source)
	if err != nil {
		log.Fatal(err.Error())
	}
}
