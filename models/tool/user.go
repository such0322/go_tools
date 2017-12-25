package tool

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"odin_tool/models/casbin"
	"strconv"
	"time"
)

type userFileds struct {
	ID        int
	Account   string
	Password  string `json:"-"`
	Name      string
	Status    int
	CreatedAt int `db:"created_at"`
}

type User struct {
	userFileds
}
type Users []User

func (m *User) DelRule(rid int) {
	e := casbin.CasbinEnforcer
	if b := e.DeleteRoleForUser(strconv.Itoa(m.ID), strconv.Itoa(rid)); !b {
		panic("角色删除失败")
	}
}

func (m *User) AddRule(rid int) {
	e := casbin.CasbinEnforcer
	if b := e.AddRoleForUser(strconv.Itoa(m.ID), strconv.Itoa(rid)); !b {
		panic("角色添加失败")
	}
}

func (m *User) GetRules() (roles []int) {
	if m.ID == 0 {
		panic(errors.New("无效的用户"))
	}
	e := casbin.CasbinEnforcer
	strRoles := e.GetRolesForUser(strconv.Itoa(m.ID))
	for _, v := range strRoles {
		r, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		roles = append(roles, r)
	}
	return
}

func (m *User) LoadByID(id int) {
	query := "select * from user where id = ?"
	row := db.QueryRowx(query, id)
	err := row.StructScan(&m.userFileds)
	if err != nil {
		panic(err)
	}
}

func (m *User) CheckPassword(password string) error {

	if m.ID == 0 {
		return errors.New("没有找到用户")
	}
	if password == "" {
		return errors.New("密码不能为空")
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	md5pw := md5Ctx.Sum(nil)
	if hex.EncodeToString(md5pw) != m.Password {
		return errors.New("密码错误")
	}
	return nil
}

func (m *User) LoadByAccount(account string) error {
	query := "select * from user where account = ?"
	row := db.QueryRowx(query, account)
	err := row.StructScan(&m.userFileds)
	if err != nil {
		return err
	}
	if m.ID == 0 {
		return errors.New("没有找到用户")
	}
	return nil
}

func (m *User) GetAll() *Users {
	rows, err := db.Queryx("select * from user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := Users{}
	for rows.Next() {
		u := User{}
		err := rows.StructScan(&u.userFileds)
		if err != nil {
			continue
		}
		users = append(users, u)
	}
	return &users
}

func (m *User) Create() {
	if m.Account == "" {
		panic(errors.New("缺少用户名"))
	}
	if m.Password == "" {
		panic(errors.New("缺少密码"))
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(m.Password))
	md5pw := md5Ctx.Sum(nil)
	m.Password = hex.EncodeToString(md5pw)
	m.Status = STATUSOPEN
	m.CreatedAt = int(time.Now().Unix())
	query := "insert into user (account, password, name, status, created_at) values (?,?,?,?,?)"
	rs, err := db.Exec(query, m.Account, m.Password, m.Name, m.Status, m.CreatedAt)
	if err != nil {
		panic(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		panic(err)
	}
	m.ID = int(id)
}

func (m *User) Close() (int64, error) {
	query := "update user set status = 0 where id = ?"
	rs, err := db.Exec(query, m.ID)
	if err != nil {
		panic(err)
	}
	return rs.RowsAffected()
}
