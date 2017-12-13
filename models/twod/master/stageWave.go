package master

import (
	"database/sql"
	"fmt"
	"odin_tool/libs"
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
	StageWaveFields
	Waves Waves
}

type StageWaves []StageWave

func (ms StageWaves) GetWaveIds() (ids []int) {
	if len(ms) == 0 {
		return ids
	}
	for _, v := range ms {
		ids = libs.AppendUniqueInt(ids, v.WaveID)
	}
	return ids
}

func (ms StageWaves) LoadWaves() Waves {
	wave := Wave{}
	if len(ms) == 0 {
		return Waves{}
	}
	wids := ms.GetWaveIds()
	waves := wave.LoadByWaveIDs(wids)
	mapper := make(map[int][]Wave)
	for _, w := range waves {
		mapper[w.WaveID] = append(mapper[w.WaveID], w)
	}
	for k, sw := range ms {
		ms[k].Waves = append(sw.Waves, mapper[sw.WaveID]...)
	}
	return waves
}

func (m *StageWave) LoadByStageID(stageID int) StageWaves {
	rows, err := db.Queryx("select * from stage_wave where stage_id = ? ", stageID)
	if err != nil {
		fmt.Println(err)
	}
	var ms StageWaves
	for rows.Next() {
		sw := StageWave{}
		if err := rows.StructScan(&sw.StageWaveFields); err != nil {
			continue
		} else {
			ms = append(ms, sw)
		}
	}
	return ms
}
