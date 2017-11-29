package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusGacha struct {
	ID    int    `json:"id"`
	Title string `json:"name"`
}
type BonusGachas struct {
	Data []BonusGacha `json:"data"`
}

func (m *BonusGachas) GetRewardNames() {
	db.Select(&m.Data, "select id, title from feature where feature_type = 'gacha'")
}
func (m *BonusGachas) GetType() string {
	return TypeGachaTicket
}

func (m *BonusGachas) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
