package bonus

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BounsItem struct {
	ID   int
	Name string
}
type BounsItems struct {
	Bonus
	Data []BounsItem
}

func (m *BounsItems) GetAll() []BounsItem {
	db.Select(&m.Data, "select id, name from item")
	fmt.Println(m.Data)
	return m.Data
}
