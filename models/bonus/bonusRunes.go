package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusRune struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusRunes struct {
	Data []BonusRune `json:"data"`
}

func (m *BonusRunes) GetRewardNames() {
	db.Select(&m.Data, "select id, name from rune")
}
func (m *BonusRunes) GetType() string {
	return TypeRune
}

func (m *BonusRunes) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
