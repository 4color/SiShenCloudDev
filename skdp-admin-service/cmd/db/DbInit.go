package db

import (
	"database/sql"
	"github.com/4color/SiShenCloudDev/skdp-admin-service/configs"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	//SqlDB, err = sql.Open(strings.ToLower(goracle.DefaultConnectionClass), config.Config.App.Db.ConnectionString)
	SqlDB, err = sql.Open("mysql", configs.Config.App.Db.ConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
