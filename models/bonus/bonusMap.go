package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusMap struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusMaps struct {
	Data []BonusMap `json:"data"`
}

func (m *BonusMaps) GetRewardNames() {
	db.Select(&m.Data, "select id, name from mission_explore")
}
func (m *BonusMaps) GetType() string {
	return TypeMap
}

func (m *BonusMaps) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
