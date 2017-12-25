package master

import (
	"database/sql"
)

type AreaFields struct {
	ID             int
	WorldID        int `db:"world_id"`
	Name           string
	Number         int
	RecommendPower int            `db:"recommend_power"`
	BgImg          string         `db:"bg_img"`
	AreaText       sql.NullString `db:"area_text"`
	StartDate      string         `db:"start_date"`
	EndDate        string         `db:"end_date"`
	InsDate        string         `db:"ins_date"`
}

type Area struct {
	AreaFields
}

type Areas []Area

func (m *Area) LoadById(id int) {
	rows := db.QueryRowx("select * from stage where id = ?", id)
	if err := rows.Err(); err != nil {
		panic(err)
	}
	rows.StructScan(&m.AreaFields)
}

func (m Area) GetByWorldID(worldID int) Areas {
	rows, err := db.Queryx("select * from area where world_id = ?", worldID)
	if err != nil {
		panic(err)
	}
	areas := Areas{}
	for rows.Next() {
		err = rows.StructScan(&m.AreaFields)
		if err != nil {
			continue
		}
		areas = append(areas, m)
	}
	return areas
}

func (m *Area) GetAll() Areas {
	areas := Areas{}
	err := db.Select(&areas, "select * from area")
	if err != nil {
		panic(err)
	}
	return areas
}
