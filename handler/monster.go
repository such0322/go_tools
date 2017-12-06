package handler

import (
	"fmt"
	"net/http"
	TM "odin_tools/models/twod/master"
	"strconv"
)

type MonsterController struct {
	Controller
}

func (c MonsterController) Test(w http.ResponseWriter, r *http.Request) {
	world := TM.NewWorld()
	world.GetAll()
	fmt.Fprintf(w, "%#v\n", world)
	fmt.Fprintf(w, "%#v\n", world.GetData())

}

func (c MonsterController) List(w http.ResponseWriter, r *http.Request) {
	stage_id, _ := strconv.Atoi(r.FormValue("stage"))
	stage := TM.Stage{}
	stage.GetById(stage_id)
	stage.LoadStageWaves().LoadWaves().LoadMonsters()

	c.data = make(map[string]interface{})
	c.data["stage"] = stage
	c.Render(w, r)
}
