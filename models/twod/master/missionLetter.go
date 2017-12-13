package master

import (
	"database/sql"
	"fmt"
	"sort"
	"time"
)

type missionLetterFields struct {
	ID                int
	Type              sql.NullString
	Priority          sql.NullInt64
	Name              string
	Description       string
	ConditionText     string         `db:"condition_text"`
	TriggerType       string         `db:"trigger_type"`
	TriggerValue      string         `db:"trigger_value"`
	CountTriggerType  sql.NullString `db:"count_trigger_type"`
	CountTriggerValue sql.NullString `db:"count_trigger_value"`
	ClearConditions   string         `db:"clear_conditions"`
	ResourceID        string         `db:"resource_id"`
	SessionFrom       sql.NullString `db:"session_from"`
	SessionTo         sql.NullString `db:"session_to"`
	PlayGamesID       sql.NullString `db:"play_games_id"`
	GameCenterID      sql.NullString `db:"game_center_id"`
	InsDate           string         `db:"ins_date"`
}

type MissionLetter struct {
	missionLetterFields
	Rewards MissionRewardLetters
}

type MissionLetters []MissionLetter

func (ms MissionLetters) Len() int {
	return len(ms)
}

func (ms MissionLetters) Less(i, j int) bool {
	return ms[i].ID > ms[j].ID
}

func (ms MissionLetters) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (m *MissionLetter) LoadByID(id int) {
	rows := db.QueryRowx("select * from mission_letter where id = ?", id)
	rows.StructScan(&m.missionLetterFields)
}

func (m *MissionLetter) GetLastDaily() MissionLetters {
	mls := MissionLetters{}
	var lastDay string
	err := db.Get(&lastDay, "select session_from from mission_letter where type = 'daily' ORDER BY session_from desc limit 1")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Queryx("select * from mission_letter where session_from = ?", lastDay)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		ml := MissionLetter{}
		err := rows.StructScan(&ml.missionLetterFields)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			mls = append(mls, ml)
		}
	}
	return mls
}

func (ms MissionLetters) GetRewardCount() (count int) {
	for _, v := range ms {
		count += len(v.Rewards)
	}
	return
}

func (ms MissionLetters) CopyAndInsert(day int, ch chan int) {
	//sql
	mlQuery := `insert into mission_letter (id, type, priority, name, description, condition_text,
	trigger_type, trigger_value, clear_conditions, resource_id, session_from, session_to, ins_date )
	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	rewardQuery := `insert into mission_reward_letter (id, mission_id, reward_type, reward_id, reward_quantity, ins_date)
	values (?, ?, ?, ?, ?, ?)`

	sort.Sort(ms)
	var sessionFrom, sessionTo time.Time
	if ms[0].SessionFrom.Valid {
		sessionFrom, _ = time.Parse("2006-01-02 15:04:05", ms[0].SessionFrom.String)
	}
	if ms[0].SessionTo.Valid {
		sessionTo, _ = time.Parse("2006-01-02 15:04:05", ms[0].SessionTo.String)
	}
	ins_date := time.Now().Format("2006-01-02 15:04:05")
	d, _ := time.ParseDuration("24h")
	//
	maxID := ms[0].ID
	sessionFrom = sessionFrom.Add(d * time.Duration(day))
	sessionTo = sessionTo.Add(d * time.Duration(day))

	maxRewardID, err := ms[0].Rewards.GetMaxID()
	if err != nil {
		fmt.Println(err)
	}
	rewardCount := ms.GetRewardCount() //13
	rcount := 0

	for k, v := range ms {
		thisID := maxID + len(ms)*(day-1) + (k + 1)
		// fmt.Println("this is :", thisID)
		_, err := db.Exec(mlQuery,
			thisID,
			v.Type,
			v.Priority,
			v.Name,
			v.Description,
			v.ConditionText,
			v.TriggerType,
			v.TriggerValue,
			v.ClearConditions,
			v.ResourceID,
			sessionFrom,
			sessionTo,
			ins_date,
		)
		if err != nil {
			panic(fmt.Sprintf("%v,%s\n", err, "mission_letter insert failed"))
		}
		for _, vv := range v.Rewards {
			rcount++
			thisRewardID := maxRewardID + rewardCount*(day-1) + rcount
			// fmt.Println("----thisRewardID:", thisID)
			_, err := db.Exec(rewardQuery,
				thisRewardID,
				thisID,
				vv.RewardType,
				vv.RewardID,
				vv.RewardQuantity,
				ins_date,
			)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	ch <- day
}

func (ms MissionLetters) LoadReward() MissionRewardLetters {
	ids := ms.GetIds()
	mlr := MissionRewardLetter{}
	mlrs := mlr.LoadByMissionIds(ids)
	mapper := make(map[int][]MissionRewardLetter)
	for _, v := range mlrs {
		mapper[v.MissionID] = append(mapper[v.MissionID], v)

	}
	for k, v := range ms {
		ms[k].Rewards = append(v.Rewards, mapper[v.ID]...)
	}
	return mlrs
}

func (ms MissionLetters) GetIds() []int {
	var ids []int
	for _, v := range ms {
		ids = append(ids, v.ID)
	}
	return ids
}
