package master

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Feature struct {
	ID          int
	StartDate   sql.NullString `db:"start_date"`
	EndDate     sql.NullString `db:"end_date"`
	OpenDate    sql.NullString `db:"open_date"`
	CloseDate   sql.NullString `db:"close_date"`
	FeatureType sql.NullString `db:"feature_type"`
	Type        sql.NullString
	Title       sql.NullString
	Description sql.NullString
	BannerUrl   sql.NullString `db:"banner_url"`
	Priority    int
	Params      sql.NullString
	InsDate     string `db:"ins_date"`
}

type Features struct{}

func (Features) GetAll() (data []Feature) {
	err := db.Select(&data, "select * from feature")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (Features) GetPage(page int) (data []Feature, pageCount int) {
	var count struct {
		Count int
	}
	db.Get(&count, "select count(*) as count from feature")
	if count.Count == 0 {
		return data, count.Count
	}

	num := 20
	offset := (page - 1) * num
	err := db.Select(&data, "select * from feature limit ?, ?", offset, num)
	if err != nil {
		log.Fatal(err)
	}
	return data, count.Count / num
}
