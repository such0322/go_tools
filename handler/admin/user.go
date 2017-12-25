package admin

import (
	"fmt"
	"net/http"
	"odin_tool/handler"
	"odin_tool/models/tool"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	handler.Controller
}

func (c UserController) DelRules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.Atoi(vars["id"])
	rid, _ := strconv.Atoi(r.FormValue("rid"))
	user := tool.User{}
	user.LoadByID(uid)
	role := tool.Role{}
	role.LoadByID(rid)
	user.DelRule(role.ID)
	http.Redirect(w, r, fmt.Sprintf("/admin/user/%d/rules", uid), 302)
}

func (c UserController) AddRules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.Atoi(vars["id"])
	rid, _ := strconv.Atoi(r.FormValue("rid"))
	user := tool.User{}
	user.LoadByID(uid)
	role := tool.Role{}
	role.LoadByID(rid)
	user.AddRule(role.ID)
	http.Redirect(w, r, fmt.Sprintf("/admin/user/%d/rules", uid), 302)
}

func (c UserController) Rules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.Atoi(vars["id"])
	user := tool.User{}
	user.LoadByID(uid)
	userRules := user.GetRules()
	fmt.Printf("%#v\n", userRules)

	role := tool.Role{}
	roles := role.GetAll()
	c.Data = make(map[string]interface{})
	c.Data["UserRules"] = userRules
	c.Data["Roles"] = roles
	c.Data["UID"] = uid
	c.Tpl = "/admin/user/rules"
	c.Render(w, r)
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/admin/user/list", http.StatusFound)
		}
	}()
	account := r.FormValue("account")
	password := r.FormValue("password")
	name := r.FormValue("name")
	user := tool.User{}
	user.Account = account
	user.Password = password
	user.Name = name
	user.Create()
	http.Redirect(w, r, "/admin/user/list", http.StatusFound)
}

func (c UserController) New(w http.ResponseWriter, r *http.Request) {
	c.Data = make(map[string]interface{})
	c.Render(w, r)
}

func (c UserController) List(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	user := tool.User{}
	users := user.GetAll()
	c.Data = make(map[string]interface{})
	c.Data["users"] = users
	c.Render(w, r)
}
