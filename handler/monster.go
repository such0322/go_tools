package handler

import (
	"fmt"
	"net/http"
	"strconv"
  TM "odin_tools/models/twod/master"
)

type MonsterController struct {
	Controller
}

func (c MonsterController) List(w http.ResponseWriter, r *http.Request) {
	var stage int
	stage, err := strconv.Atoi(r.FormValue("stage"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stage)
  monsters := TM.

	c.data = make(map[string]interface{})

	c.Render(w, r)
}
