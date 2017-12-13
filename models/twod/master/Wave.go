package master

import (
	"fmt"
	"odin_tool/libs"

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
	WaveFields
	Ml MonsterLevel
}

type Waves []Wave

func (m *Wave) LoadByWaveIDs(ids []int) Waves {
	query := "select * from wave where wave_id in (?)"
	query, args, err := sqlx.In(query, ids)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Queryx(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	ms := Waves{}
	for rows.Next() {
		w := Wave{}
		if err := rows.StructScan(&w.WaveFields); err != nil {
			continue
		} else {
			ms = append(ms, w)
		}
	}
	return ms
}

func (ms Waves) LoadMonsters() MonsterLevels {
	if len(ms) == 0 {
		return MonsterLevels{}
	}
	mlIDs := ms.GetMonsterLevelIDs()
	ml := MonsterLevel{}
	mls := ml.LoadByIDs(mlIDs)
	mapper := make(map[int]MonsterLevel)
	for _, v := range mls {
		mapper[v.ID] = v
	}
	for _, v := range ms {
		v.Ml = mapper[v.MonsterLevelID]
	}
	return mls
}

func (ms Waves) GetMonsterLevelIDs() (ids []int) {
	len := len(ms)
	if len == 0 {
		return ids
	}
	for _, v := range ms {
		ids = libs.AppendUniqueInt(ids, v.MonsterLevelID)
	}
	return ids
}
