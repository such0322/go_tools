package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusItems struct {
	Data []BonusItem `json:"data"`
}

func (m *BonusItems) GetRewardNames() {
	db.Select(&m.Data, "select id, name from item")
}
func (m *BonusItems) GetType() string {
	return TypeItem
}

func (m *BonusItems) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
