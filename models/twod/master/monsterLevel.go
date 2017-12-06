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
	data MonsterLevelFields
}
type MonsterLevels struct {
	list []*MonsterLevel
}

func (ms *MonsterLevels) LoadByIDs(ids []int) {
	query := "select * from monster_level where id in (?)"
	query, args, err := sqlx.In(query, ids)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Queryx(query, args...)
	for rows.Next() {
		ml := new(MonsterLevel)
		if err := rows.StructScan(&ml.data); err != nil {
			fmt.Println(err)
		} else {
			ms.list = append(ms.list, ml)
		}
	}
}

// func (ms *MonsterLevels) GetPage(page int, prePage int, url string, where string) template.HTML {
// 	if page <= 0 {
// 		page = 1
// 	}
// 	var count int
// 	if err := db.Get(&count, fmt.Sprintf("select count(*) from %s %s", TABLE, where)); err != nil {
// 		fmt.Println(err)
// 	}
// 	pages := &libs.Pages{Count: count, Page: page, PrePage: prePage, Url: url}
// 	offset := (page - 1) * prePage
// 	if err := db.Select(&(ms.data), fmt.Sprintf("select * from %s %s limit ?, ?", TABLE, where), offset, prePage); err != nil {
// 		fmt.Println(err)
// 	}
// 	return template.HTML(pages.Get())
// }
