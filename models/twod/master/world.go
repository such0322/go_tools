package master

import (
	"odin_tools/models"

	_ "github.com/go-sql-driver/mysql"
)

type World struct {
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

type Worlds struct {
	data []World
	models.Model
}

func NewWorld() *Worlds {
	var worlds = new(Worlds)
	worlds.SetTable("world")
	worlds.SetData(worlds.data)
	worlds.SetSchema(World{})
	return worlds
}
func (m *Worlds) GetData() []World {
	return m.data
}
