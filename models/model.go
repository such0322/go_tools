package models

import (
	"fmt"
	"reflect"
	"strings"
)

type Modeler interface {
	SetTable(table string)
	// Data() []interface{}
}

type Model struct {
	table  string
	schema interface{}
	data   interface{}
}

func (m *Model) SetTable(table string) {
	m.table = table
}

func (m *Model) SetSchema(schema interface{}) {
	m.schema = schema
}

//
func (m *Model) SetData(data interface{}) {
	//TODO 把m.data ptr = data ptr
	// fmt.Printf("%#v\n", data)

	// vp := reflect.ValueOf(&data)
	// v := reflect.Indirect(vp)
	// fmt.Printf("%#v %+v\n", vp, vp)
	// fmt.Printf("%#v %+v\n", v, v)

	// m.data = reflect.New(v.Type())
	// m.data = vp.Interface()
	m.data = &data
	fmt.Printf("%#v %+v\n", m.data, m.data)

}

func fieldsTypes(refType reflect.Type) map[string]reflect.Type {
	flen := refType.NumField()
	var rs = make(map[string]reflect.Type, flen)
	for i := 0; i < flen; i++ {
		field := refType.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			name := strings.ToLower(field.Name)
			rs[name] = field.Type
		} else {
			rs[tag] = field.Type
		}
	}
	return rs
}

func (m *Model) GetAll() interface{} {
	refVal := reflect.ValueOf(m.schema)
	if refVal.Kind() != reflect.Struct {
		fmt.Println("not struct")
	}
	base := refVal.Type()
	slice := reflect.MakeSlice(reflect.SliceOf(base), 0, 10)
	rows, err := db.Query(fmt.Sprintf("select * from %s", m.table))
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	var values = make([]interface{}, len(cols))
	for rows.Next() {
		vp := reflect.New(base)
		v := reflect.Indirect(vp)
		for i := 0; i < base.NumField(); i++ {
			f := reflect.Indirect(v).Field(i)
			values[i] = f.Addr().Interface()
		}
		err := rows.Scan(values...)
		if err != nil {
			fmt.Println(err)
		}
		slice = reflect.Append(slice, v)
	}
	// fmt.Printf("getall: %#v\n %+v\n", m.data, m.data)
	m.data = slice.Interface()

	return slice.Interface()
}
