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
	Areas Areas
}

type Worlds []World

func (m *World) LoadByID(id int) {
	rows := db.QueryRowx("select * from world where id = ?", id)
	if err := rows.Err(); err != nil {
		panic(err)
	}
	rows.StructScan(&m.WorldFields)
}

func (m *World) LoadAreas() Areas {
	if m.ID == 0 {
		panic(errors.New("LoadAreas没有找到World;"))
	}
	area := Area{}
	m.Areas = area.GetByWorldID(m.ID)
	return m.Areas
}

func (m *World) GetAll() Worlds {
	rows, err := db.Queryx("select * from world")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	worlds := Worlds{}
	for rows.Next() {
		w := World{}
		err := rows.StructScan(&w.WorldFields)
		if err != nil {
			continue
		}
		worlds = append(worlds, w)
	}
	return worlds
}
