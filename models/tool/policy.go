package tool

import "errors"

const POLICYSTATUSPUBLIC = 2

type PolicyFields struct {
	ID        int
	Name      string
	Method    string
	Pattern   string
	Group     string
	Status    int
	CreatedAt int `db:"created_at"`
}

type Policy struct {
	PolicyFields
}

type Policies []Policy

type PolicyMap map[string]map[string]Policy

var pmapper PolicyMap

func (m *Policy) GetPMapper() PolicyMap {
	if pmapper == nil {
		ps := m.GetAll()
		pmapper = make(map[string]map[string]Policy)
		for _, p := range ps {
			if _, ok := pmapper[p.Name]; !ok {
				pmapper[p.Name] = make(map[string]Policy)
			}
			pmapper[p.Name][p.Method] = p
		}
	}
	return pmapper
}

func (m *Policy) GetAll() Policies {
	query := "select * from policy"
	rows, err := db.Queryx(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	ps := Policies{}
	for rows.Next() {
		p := Policy{}
		err := rows.StructScan(&p.PolicyFields)
		if err != nil {
			continue
		}
		ps = append(ps, p)
	}
	return ps
}

func (m *Policy) StatusPublic() bool {
	return m.Status == POLICYSTATUSPUBLIC
}
func (m *Policy) StatusOpen() bool {
	return m.Status != STATUSCLOSE
}

func (m *Policy) LoadByNameMethod(name, method string) {
	pmapper := m.GetPMapper()
	p, ok := pmapper[name][method]
	if !ok {
		panic(errors.New("错误的链接"))
	}
	m.PolicyFields = p.PolicyFields
}
