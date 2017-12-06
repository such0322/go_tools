package master

import (
	"fmt"
	"odin_tools/libs"

	"github.com/jmoiron/sqlx"
)

type WaveFields struct {
	ID             int
	WaveID         int    `db:"wave_id"`
	OrderID        int    `db:"order_id"`
	MonsterLevelID int    `db:"monster_level_id"`
	InsDate        string `db:"ins_date"`
}

type Wave struct {
	data WaveFields
	Ml   *MonsterLevel
}

type Waves struct {
	list []*Wave
}

func (ms *Waves) LoadByWaveIDs(ids []int) {
	query := "select * from wave where wave_id in (?)"
	query, args, err := sqlx.In(query, ids)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Queryx(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		wave := new(Wave)
		if err := rows.StructScan(&wave.data); err != nil {
			fmt.Println(err)
		} else {
			ms.list = append(ms.list, wave)
		}
	}
}

func (ms *Waves) LoadMonsters() *MonsterLevels {
	ml := new(MonsterLevels)
	if len(ms.list) == 0 {
		return ml
	}
	mlIDs := ms.GetMonsterLevelIDs()
	ml.LoadByIDs(mlIDs)
	mapper := make(map[int]*MonsterLevel)
	for _, v := range ml.list {
		mapper[v.data.ID] = v
	}
	for _, v := range ms.list {
		v.Ml = mapper[v.data.MonsterLevelID]
	}
	return ml
}

func (ms *Waves) GetMonsterLevelIDs() (ids []int) {
	len := len(ms.list)
	if len == 0 {
		return ids
	}
	for _, v := range ms.list {
		ids = libs.AppendUniqueInt(ids, v.data.MonsterLevelID)
	}
	return ids
}

func (ms *Waves) GetList() []*Wave {
	return ms.list
}

func (ms *Waves) SetList(waves ...*Wave) {
	for _, w := range waves {
		ms.list = append(ms.list, w)
	}

}
