package models

import (
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Package struct {
	Type string `json:"type"`
	ID   int64  `json:"id"`
	Num  int64  `json:"num"`
}

type Ttt struct {
	Id      int64     `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	Package []Package `db:"package" json:"package"`
}

func (t *Ttt) ToJson() string {
	rs, _ := json.Marshal(t)
	return string(rs)
}

func conn() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:123456@(127.0.0.1:33306)/odin_test?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	// defer db.Close()
	return db
}
func (t *Ttt) GetAll() {
	db := conn()
	ts := []Ttt{}
	db.Select(&ts, "select * from ttt")
	fmt.Println(ts)
}

func (t *Ttt) Create() {
	db := conn()
	pkg, _ := json.Marshal(t.Package)
	rs, _ := db.Exec("insert into ttt (name,package) values (?,?)", t.Name, pkg)
	id, _ := rs.LastInsertId()
	t.Id = id

}
func (t *Ttt) Get(id int64) {
	db := conn()
	defer db.Close()
	row := db.QueryRowx("select * from ttt where id = ?", id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
	}
	var name string
	var pkg string
	row.Scan(&id, &name, &pkg)
	t.Id = id
	t.Name = name
	err := json.Unmarshal([]byte(pkg), &t.Package)
	if err != nil {
		fmt.Println(err)
	}
}
