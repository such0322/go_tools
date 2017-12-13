package master

import (
	"errors"
)

type WorldFields struct {
	ID                          int
	Name                        string
	Chapter                     string
	UnlockCondition             string `db:"unlock_condition"`
	ContentsLockUnlockCondition string `db:"contents_lock_unlock_condition"`
	BgImg                       string `db:"bg_img"`
	SlotDiff                    int    `db:"slot_diff"`
	StartDate                   string `db:"start_date"`
	EndDate                     string `db:"end_date"`
	InsDate                     string `db:"ins_date"`
}
type World struct {
	WorldFields
	Areas *Areas
}

type Worlds []World

func (m *World) GetByID(id int) *World {
	rows := db.QueryRowx("select * from world where id = ?", id)
	if err := rows.Err(); err != nil {
		panic(err)
	}
	rows.StructScan(&m.WorldFields)
	return m
}

func (m *World) LoadAreas() *Areas {
	if m.ID == 0 {
		panic(errors.New("LoadAreas没有找到World;"))
	}
	area := Area{}
	m.Areas = area.LoadByWorldID(m.ID)
	return m.Areas
}

func (m World) LoadAll() *Worlds {
	rows, err := db.Queryx("select * from world")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	worlds := Worlds{}
	for rows.Next() {
		err = rows.StructScan(&m.WorldFields)
		if err != nil {
			continue
		}
		worlds = append(worlds, m)
	}
	return &worlds
}
