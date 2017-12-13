package master

import (
	"database/sql"
	"fmt"
	"html/template"
	"odin_tool/libs"
	"sort"
)

//StageFields 表字段
type stageFields struct {
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

type Stage struct {
	stageFields
	StageWaves StageWaves
}

type Stages []Stage

type StageMonsters struct {
	Order  int
	WaveID int
	IsBoss bool
	MLs    MonsterLevels
}
type SMSlice []StageMonsters

func (m SMSlice) Len() int {
	return len(m)
}
func (m SMSlice) Less(i, j int) bool {
	if m[i].Order < m[j].Order {
		return true
	} else if m[i].Order > m[j].Order {
		return false
	} else {
		return m[i].WaveID < m[j].WaveID
	}

}
func (m SMSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *Stage) GetOrderMonsters() []StageMonsters {
	sw := m.StageWaves
	sms := SMSlice{}
	for _, v := range sw {
		sm := StageMonsters{}
		sm.IsBoss = (v.IsBoss != 0)
		sm.Order = v.OrderID
		sm.WaveID = v.WaveID
		sm.MLs = sw.LoadWaves().LoadMonsters()
		sms = append(sms, sm)
		sort.Sort(sms)
	}
	return sms
}

func (m *Stage) LoadById(id int) {
	rows := db.QueryRowx("select * from stage where id = ?", id)
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}
	rows.StructScan(&m.stageFields)
}

//LoadStageWaves 获取对应stageWaves
func (m *Stage) LoadStageWaves() StageWaves {
	if m.ID == 0 {
		return StageWaves{}
	}
	stageWave := StageWave{}
	stageWaves := stageWave.LoadByStageID(m.ID)
	m.StageWaves = stageWaves
	return m.StageWaves
}

func (m Stage) GetPage(page int, prePage int, url string, where string, args ...interface{}) (*Stages, template.HTML) {
	if page <= 0 {
		page = 1
	}
	var cquery, query string
	if where == "" {
		cquery = "select count(*) from stage"
		query = "select * from stage limit ?, ?"
	} else {
		cquery = fmt.Sprintf("select count(*) from stage where %s", where)
		query = fmt.Sprintf(`select * from stage where  %s limit ?, ?`, where)
	}
	var count int
	if err := db.Get(&count, cquery, args...); err != nil {
		fmt.Println(err)
	}
	pages := &libs.Pages{Count: count, Page: page, PrePage: prePage, Url: url}
	offset := (page - 1) * prePage

	args = append(args, offset, prePage)
	rows, err := db.Queryx(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	stages := Stages{}
	for rows.Next() {
		err := rows.StructScan(&m.stageFields)
		if err != nil {
			continue
		}
		stages = append(stages, m)
	}
	return &stages, pages.Get()
}
