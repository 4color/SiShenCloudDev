package db

import (
	"github.com/4color/SiShenCloudDev/skdp-admin-service/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var SqlDB *gorm.DB

func init() {
	var err error
	//SqlDB, err = sql.Open(strings.ToLower(goracle.DefaultConnectionClass), config.Config.App.Db.ConnectionString)
	SqlDB, err = gorm.Open("mysql", configs.Config.App.Db.ConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}

	//defer SqlDB.Close()

	err = SqlDB.DB().Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
