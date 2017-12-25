package master

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:123456@tcp(192.168.7.120:3306)/twod_enish_master_2?charset=utf8")
	//限制sql连接数，否则协程会报连接数问题
	db.SetMaxOpenConns(100)
	// db.SetMaxIdleConns(50)
	// db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}
