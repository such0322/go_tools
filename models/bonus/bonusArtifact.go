package bonus

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BonusArtifact struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type BonusArtifacts struct {
	Data []BonusArtifact `json:"data"`
}

func (m *BonusArtifacts) GetRewardNames() {
	db.Select(&m.Data, "select id, name from artifact")
}
func (m *BonusArtifacts) GetType() string {
	return TypeArtifact
}

func (m *BonusArtifacts) ToJson() string {
	json, err := json.Marshal(m.Data)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
