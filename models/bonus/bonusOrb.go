package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusOrb struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusOrbs struct {
	Data []BonusOrb `json:"data"`
}

func (m *BonusOrbs) GetRewardNames() {
	db.Select(&m.Data, "select id, name from promotion_orb")
}
func (m *BonusOrbs) GetType() string {
	return TypeORB
}

func (m *BonusOrbs) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
