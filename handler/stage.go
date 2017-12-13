package handler

import (
	"fmt"
	"net/http"
	TM "odin_tool/models/twod/master"
	"strconv"

	"github.com/gorilla/mux"
)

type StageController struct {
	Controller
}

func (c StageController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	stage := TM.Stage{}
	stage.LoadById(id)
	stage.LoadStageWaves().LoadWaves().LoadMonsters()

	c.data = make(map[string]interface{})
	c.data["stage"] = stage
	c.data["orderMonsters"] = stage.GetOrderMonsters()
	c.tpl = "/stage/get"
	c.Render(w, r)

}

func (c StageController) List(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	p, err := strconv.Atoi(r.FormValue("p"))
	if err != nil {
		p = 1
	}
	url := "/stage/list"
	where := ""
	var args []interface{}
	areaID, err := strconv.Atoi(r.FormValue("area_id"))
	if err == nil {
		url = "/stage/list?area_id=" + strconv.Itoa(areaID)
		where = "area_id = ? "
		args = append(args, areaID)
	}

	stage := TM.Stage{}
	stages, pager := stage.GetPage(p, 100, url, where, args...)
	c.data = make(map[string]interface{})
	c.data["Pager"] = pager
	c.data["Stages"] = stages
	c.Render(w, r)
}
