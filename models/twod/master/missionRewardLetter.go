package master

import (
	"errors"
	"fmt"
	"sort"

	"github.com/jmoiron/sqlx"
)

type missionRewardLetterFields struct {
	ID             int
	MissionID      int    `db:"mission_id"`
	RewardType     string `db:"reward_type"`
	RewardID       int    `db:"reward_id"`
	RewardQuantity int    `db:"reward_quantity"`
	InsDate        string `db:"ins_date"`
}

type MissionRewardLetter struct {
	missionRewardLetterFields
}
type MissionRewardLetters []MissionRewardLetter

func (ms MissionRewardLetters) Len() int {
	return len(ms)
}

func (ms MissionRewardLetters) Less(i, j int) bool {
	return ms[i].ID > ms[j].ID
}

func (ms MissionRewardLetters) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (m *MissionRewardLetter) LoadByMissionIds(ids []int) MissionRewardLetters {
	query := "select * from mission_reward_letter where mission_id in (?)"
	query, args, err := sqlx.In(query, ids)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Queryx(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	mrls := MissionRewardLetters{}
	for rows.Next() {
		mrl := MissionRewardLetter{}
		err := rows.StructScan(&mrl.missionRewardLetterFields)
		if err != nil {
			fmt.Println(err)
			continue
		}
		mrls = append(mrls, mrl)
	}
	return mrls
}

func (ms MissionRewardLetters) GetMaxID() (int, error) {
	if len(ms) == 0 {
		return 0, errors.New("日常任务缺少奖励数据")
	}
	sort.Sort(ms)
	return ms[0].ID, nil
}
