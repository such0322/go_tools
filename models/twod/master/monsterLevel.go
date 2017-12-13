package master

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MonsterLevelFields struct {
	ID                        int
	MonsterID                 int `db:"monster_id"`
	MonsterLevel              int `db:"monster_level"`
	Hp                        sql.NullInt64
	Exp                       int
	Lupi                      int
	MonsterRewardGroupID      int            `db:"monster_reward_group_id"`
	MonsterRewardRarityWeight sql.NullString `db:"monster_reward_rarity_weight"`
	Strength                  sql.NullInt64
	JobID                     sql.NullInt64 `db:"job_id"`
	Attachments               string
	Appear                    string
	Skill1                    int    `db:"skill_1"`
	Skill2                    int    `db:"skill_2"`
	Skill3                    int    `db:"skill_3"`
	Skill4                    int    `db:"skill_4"`
	Skill5                    int    `db:"skill_5"`
	Skill6                    int    `db:"skill_6"`
	Skill7                    int    `db:"skill_7"`
	Skill8                    int    `db:"skill_8"`
	Skill9                    int    `db:"skill_9"`
	Skill10                   int    `db:"skill_10"`
	InsDate                   string `db:"ins_date"`
}
type MonsterLevel struct {
	MonsterLevelFields
}
type MonsterLevels []MonsterLevel

func (m *MonsterLevel) LoadByIDs(ids []int) MonsterLevels {
	query := "select * from monster_level where id in (?)"
	query, args, err := sqlx.In(query, ids)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Queryx(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	mls := MonsterLevels{}
	for rows.Next() {
		ml := MonsterLevel{}
		if err := rows.StructScan(&ml.MonsterLevelFields); err != nil {
			continue
		} else {
			mls = append(mls, ml)
		}
	}
	return mls
}
