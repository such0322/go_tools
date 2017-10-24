package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Ttt struct {
	Id          int
	Name        string
	DeviceModel sql.NullString `db:"device_model"`
	InsDate     string         `db:"ins_date"`
}

func (t *Ttt) GetAll() {
	db, err := sqlx.Connect("mysql", "root:123456@(127.0.0.1:33306)/bridge_enish_user?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	ts := []Ttt{}
	db.Select(&ts, "select * from ttt")
	fmt.Println(ts)
}
