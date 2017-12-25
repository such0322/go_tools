package libs

import (
	"net/http"
	mc "odin_tool/models/casbin"
	"odin_tool/models/tool"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const SUPERADMIN = 1

func Casbin(inner http.Handler, routeName, routeMethod string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		policy := tool.Policy{}
		policy.LoadByNameMethod(routeName, routeMethod)
		if !policy.StatusOpen() {
			http.Redirect(w, r, "/error", 302)
		}
		if policy.StatusPublic() {
			inner.ServeHTTP(w, r)
			return
		}

		//进行权限验证
		currentUser := CurrentUser
		if currentUser.ID == 0 {
			http.Redirect(w, r, "/auth/login", 302)
			return
		}
		ce := mc.CasbinEnforcer
		roles := ce.GetRolesForUser(strconv.Itoa(currentUser.ID))
		for _, role := range roles {
			if strconv.Itoa(SUPERADMIN) == role {
				inner.ServeHTTP(w, r)
				return
			}
			if ce.Enforce(role, strconv.Itoa(policy.ID), routeMethod) {
				inner.ServeHTTP(w, r)
				return
			}
		}
		http.Redirect(w, r, "/error", 302)

	})
}
