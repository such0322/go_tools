package casbin

import (
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

const SUPERADMIN = 1

var CasbinEnforcer *casbin.Enforcer

func init() {
	a := xormadapter.NewAdapter("mysql", "root:123456@tcp(192.168.7.120:3306)/odin_tool", true)
	CasbinEnforcer = casbin.NewEnforcer("conf/casbin/rbac_model.conf", a)
	CasbinEnforcer.EnableLog(false)
}
