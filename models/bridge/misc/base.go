package misc

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:123456@tcp(192.168.7.120:3306)/bridge_enish_misc?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
}
