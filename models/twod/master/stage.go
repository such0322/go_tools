package master

import (
	"database/sql"
	"fmt"
	"html/template"
	"odin_tools/libs"
)

type Stage struct {
	ID               int
	AreaID           sql.NullInt64 `db:"area_id"`
	Grade            int
	ScenarioID       sql.NullString `db:"scenario_id"`
	BgImg            string         `db:"bg_img"`
	UnlockCondition  sql.NullInt64  `db:"unlock_condition"`
	Ap               int
	Exp              int
	Lupi             int
	Mode             int
	CanClosed        string `db:"can_closed"`
	NpcNum           int    `db:"npc_num"`
	ParticipantLimit int    `db:"participant_limit"`
	RecommendedLevel int    `db:"recommended_level"`
	MaxRank          int    `db:"max_rank"`
	PlasmaMode       int    `db:"plasma_mode"` //大概是没用的字段
	FriendDiff       int    `db:"friend_diff"`
	CostDiff         int    `db:"cost_diff"`
	ApDiff           int    `db:"ap_diff"`
	Artifact         sql.NullInt64
	ItemBundleSet    sql.NullString `db:"item_bundle_set"`
	Name             string
	Number           int
	Description      string
	LandscapeID      sql.NullInt64  `db:"landscape_id"`
	NormalBgm        sql.NullString `db:"normal_bgm"`
	BossBgm          sql.NullString `db:"boss_bgm"`
	ContinueCount    sql.NullInt64  `db:"continue_count"`
	ContinueOrb      int            `db:"continue_orb"`
	IsInfinity       sql.NullInt64  `db:"is_infinity"`
	Abilities        string
	OpenDate         string `db:"open_date"`
	InsDate          string `db:"ins_date"`
}

type Stages struct {
	Data []Stage
}

func (ms *Stages) GetAll() {
	err := db.Select(ms.Data, "select * from stage")
	if err != nil {
		fmt.Println(err)
	}
}

func (ms *Stages) GetPage(page int, prePage int, url string) (pager template.HTML) {
	if page <= 0 {
		page = 1
	}
	var count int
	if err := db.Get(&count, "select count(*) from stage"); err != nil {
		fmt.Println(err)
	}
	pages := &libs.Pages{Count: count, Page: page, PrePage: prePage, Url: url}
	offset := (page - 1) * prePage
	if err := db.Select(&(ms.Data), "select * from stage  limit ?, ?", offset, prePage); err != nil {
		fmt.Println(err)
	}
	return template.HTML(pages.Get())
}
