package handler

import (
	"net/http"
	TM "odin_tools/models/twod/master"
	"strconv"
)

type StageController struct {
	Controller
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
	c.data["Stages"] = stages
	c.Render(w, r)
}
