package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusItemBundle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusItemBundles struct {
	Data []BonusItemBundle `json:"data"`
}

func (m *BonusItemBundles) GetRewardNames() {
	db.Select(&m.Data, "select id, name from item_bundle")
}
func (m *BonusItemBundles) GetType() string {
	return TypeItemBundle
}

func (m *BonusItemBundles) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
