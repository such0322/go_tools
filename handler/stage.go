package handler

import (
	"net/http"
	TM "odin_tools/models/twod/master"
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
	stage.GetById(id)
	stage.LoadStageWaves().LoadWaves().LoadMonsters()

	c.data = make(map[string]interface{})
	c.data["stage"] = stage.GetData()
	c.data["orderMonsters"] = stage.GetOrderMonsters()
	c.tpl = "/stage/get"
	c.Render(w, r)

}

func (c StageController) List(w http.ResponseWriter, r *http.Request) {
	var p int
	if r.FormValue("p") == "" {
		p = 1
	} else {
		var err error
		p, err = strconv.Atoi(r.FormValue("p"))
		if err != nil {
			p = 1
		}
	}
	stages := TM.Stages{}
	pager := stages.GetPage(p, 100, "/stage/list")
	c.data = make(map[string]interface{})
	c.data["Pager"] = pager
	c.data["Stages"] = stages.GetData()
	c.Render(w, r)
}
