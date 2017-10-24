package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DeviceCredential struct {
	Id            int    `db:"id"`
	AppId         string `db:"app_id"`
	Platform      string
	Store         string
	Status        string
	Uiid          string
	Udid          string
	UdidType      string `db:"udid_type"`
	Guid          string
	Uid           int
	DeviceModel   sql.NullString `db:"device_model"`
	DeviceVersion sql.NullString `db:"device_version"`
	Uaid          sql.NullString
	UaidEnabled   sql.NullString `db:"uaid_enabled"`
	InsDate       string         `db:"ins_date"`
	LastModDate   string         `db:"last_mod_date"`
	Grade         int
	Lang          sql.NullString
	Country       sql.NullString
}

func (dc *DeviceCredential) GetByGuid() {

	db, err := sqlx.Connect("mysql", "root:123456@(127.0.0.1:33306)/bridge_enish_user?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	dcs := []DeviceCredential{}
	// tables := [...]string{"aa", "bbb"}
	// db.Select(tables, "show tables")
	// fmt.Println(tables)

	db.Select(&dcs, "select * from device_credentials where id = 5")
	//
	fmt.Println(dcs)
}
