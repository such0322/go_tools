package tool

type RoleField struct {
	ID        int
	Name      string
	Status    int
	CreatedAt int `db:"created_at"`
}

type Role struct {
	RoleField
}
type Roles []Role

func (m *Role) LoadByID(id int) {
	query := "select * from role where id = ?"
	row := db.QueryRowx(query, id)
	err := row.StructScan(&m.RoleField)
	if err != nil {
		panic(err)
	}
}

func (m *Role) GetAll() Roles {
	query := "select * from role "
	rows, err := db.Queryx(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	roles := Roles{}
	for rows.Next() {
		r := Role{}
		err := rows.StructScan(&r.RoleField)
		if err != nil {
			continue
		}
		roles = append(roles, r)
	}
	return roles
}
