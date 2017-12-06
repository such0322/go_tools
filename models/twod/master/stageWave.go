package master

import (
	"database/sql"
	"fmt"
	"odin_tools/libs"
)

type StageWaveFields struct {
	ID        int
	StageID   int `db:"stage_id"`
	OrderID   int `db:"order_id"`
	WaveID    int `db:"wave_id"`
	Weight    int
	Camera    sql.NullString
	IsBoss    int            `db:"is_boss"`
	OpenDate  sql.NullString `db:"open_date"`
	CloseDate sql.NullString `db:"close_date"`
}

type StageWave struct {
	data  StageWaveFields
	Waves *Waves
}

func (m *StageWave) GetData() StageWaveFields {
	return m.data
}

type StageWaves struct {
	list []*StageWave
}

func (ms *StageWaves) GetWaveIds() (ids []int) {
	len := len(ms.list)
	if len == 0 {
		return ids
	}
	for _, v := range ms.list {
		ids = libs.AppendUniqueInt(ids, v.data.WaveID)
	}
	return ids
}

func (ms *StageWaves) LoadWaves() *Waves {
	waves := new(Waves)
	if len(ms.list) == 0 {
		return waves
	}
	wids := ms.GetWaveIds()
	waves.LoadByWaveIDs(wids)
	mapper := make(map[int][]*Wave)
	for _, w := range waves.list {
		mapper[w.data.WaveID] = append(mapper[w.data.WaveID], w)
	}
	for _, sw := range ms.list {
		w := new(Waves)
		w.SetList(mapper[sw.data.WaveID]...)
		sw.Waves = w
	}
	return waves
}

// func (ms *StageWaves) LoadMonsters() *StageWaves {
// 	fmt.Println("load Monsters")
// 	if len(ms.list) == 0 {
// 		//TODO panic
// 		return ms
// 	}
//
// 	for _, v := range ms.list {
// 		fmt.Println(v)
// 	}
//
// 	return ms
// }

func (ms *StageWaves) LoadByStageID(stageID int) {
	rows, err := db.Queryx("select * from stage_wave where stage_id = ? ", stageID)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		sw := new(StageWave)
		if err := rows.StructScan(&sw.data); err != nil {
			fmt.Println(err)
		} else {
			ms.list = append(ms.list, sw)
		}
	}
}

func (ms *StageWaves) GetList() []*StageWave {
	return ms.list
}
