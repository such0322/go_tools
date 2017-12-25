package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"odin_tool/libs"
	"odin_tool/models/tool"
)

type AuthController struct {
	Controller
}

func (c Controller) Logout(w http.ResponseWriter, r *http.Request) {
	libs.Session.Delete("currentUser")
	libs.CurrentUser = tool.User{}
	http.Redirect(w, r, "/", 302)
}

func (c Controller) DoLogin(w http.ResponseWriter, r *http.Request) {
	account := r.FormValue("account")
	password := r.FormValue("password")
	user := tool.User{}
	if err := user.LoadByAccount(account); err != nil {
		fmt.Println(err)
	}
	if err := user.CheckPassword(password); err != nil {
		fmt.Println(err)
	}
	//
	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	libs.Session.Set("currentUser", string(jsonUser))
	http.Redirect(w, r, "/", 302)
}

func (c Controller) Login(w http.ResponseWriter, r *http.Request) {
	c.Data = make(map[string]interface{})
	c.Render(w, r)
}
