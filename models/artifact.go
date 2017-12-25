package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Artifact struct {
	ID                       int
	Name                     string
	FemaleName               string `db:"female_name"`
	Description              sql.NullString
	Type                     string
	Subtype                  sql.NullString
	Animatype                string
	AutoattackSkillId        sql.NullInt64  `db:"autoattack_skill_id"`
	PassiveSkillIds          sql.NullString `db:"passive_skill_ids"`
	ActiveSkillIds           sql.NullString `db:"active_skill_ids"`
	Position                 sql.NullString
	Gender                   sql.NullString
	Element                  sql.NullString
	JobIds                   string `db:"job_ids"`
	Icon                     string
	Model                    string
	Texture                  string
	Resource                 sql.NullString
	ViewParam                string `db:"view_param"`
	Motion                   string
	GrowthType               string         `db:"growth_type"`
	EvolveInfo               sql.NullString `db:"evolve_info"`
	Buster                   string
	Rarity                   int
	LevelMax                 int            `db:"level_max"`
	RuneStabilityNum         int            `db:"rune_stability_num"`
	RuneSlotNum              int            `db:"rune_slot_num"`
	RuneLupi                 sql.NullInt64  `db:"rune_lupi"`
	RunePotentialGroupId     sql.NullInt64  `db:"rune_potential_group_id"`
	RuneEquipmentGroupId     sql.NullInt64  `db:"rune_equipment_group_id"`
	RuneUnexpectedGroupId    int            `db:"rune_unexpected_group_id"`
	ArtifactSynthesisGroupId sql.NullInt64  `db:"artifact_synthesis_group_id"`
	TranscendenceIds         sql.NullString `db:"transcendence_ids"`
	HpMin                    sql.NullInt64  `db:"hp_min"`
	HpMax                    sql.NullInt64  `db:"hp_max"`
	MpMin                    int            `db:"mp_min"`
	MpMax                    int            `db:"mp_max"`
	AtkPMin                  sql.NullInt64  `db:"atk_p_min"`
	AtkPMax                  sql.NullInt64  `db:"atk_p_max"`
	AtkMMin                  sql.NullInt64  `db:"atk_m_min"`
	AtkMMax                  sql.NullInt64  `db:"atk_m_max"`
	AidMMin                  int            `db:"aid_m_min"`
	AidMMax                  int            `db:"aid_m_max"`
	DefPMin                  sql.NullInt64  `db:"def_p_min"`
	DefPMax                  sql.NullInt64  `db:"def_p_max"`
	DefMMin                  sql.NullInt64  `db:"def_m_min"`
	DefMMax                  sql.NullInt64  `db:"def_m_max"`
	SpeedMin                 sql.NullInt64  `db:"speed_min"`
	SpeedMax                 sql.NullInt64  `db:"speed_max"`
	AccuracyMin              sql.NullInt64  `db:"accuracy_min"`
	AccuracyMax              sql.NullInt64  `db:"accuracy_max"`
	DodgeMin                 sql.NullInt64  `db:"dodge_min"`
	DodgeMax                 sql.NullInt64  `db:"dodge_max"`
	CriticalMin              sql.NullInt64  `db:"critical_min"`
	CriticalMax              sql.NullInt64  `db:"critical_max"`
	Cost                     sql.NullInt64
	SeriesId                 sql.NullInt64 `db:"series_id"`
	ImportantFlag            sql.NullInt64 `db:"important_flag"`
	IsLegend                 sql.NullInt64 `db:"is_legend"`
	EvolveSynthesisNum       sql.NullInt64 `db:"evolve_synthesis_num"`
	IsDual                   sql.NullInt64 `db:"is_dual"`
	OpenDate                 string        `db:"open_date"`
	InsDate                  string        `db:"ins_date"`
}

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("mysql", "root:123456@tcp(192.168.7.120:3306)/twod_enish_master_2?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
}

func (a *Artifact) LoadById(id int) {
	db.Get(a, "select * from artifact where id = ?", id)
}

type Artifacts struct{}

func (Artifacts) GetAll() (as []Artifact) {
	db.Select(&as, "select * from artifact")
	return as
}
