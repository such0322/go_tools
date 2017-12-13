package handler

import (
	"fmt"
	"net/http"
	TM "odin_tool/models/twod/master"
	"strconv"
)

type MonsterController struct {
	Controller
}

func (c MonsterController) Test(w http.ResponseWriter, r *http.Request) {

}

func (c MonsterController) List(w http.ResponseWriter, r *http.Request) {
	stage_id, _ := strconv.Atoi(r.FormValue("stage"))
	stage := TM.Stage{}
	stage.LoadById(stage_id)
	stage.LoadStageWaves().LoadWaves().LoadMonsters()

	fmt.Fprintf(w, "%#v\n", stage)
	c.data = make(map[string]interface{})
	c.data["stage"] = stage
	// c.Render(w, r)
}
