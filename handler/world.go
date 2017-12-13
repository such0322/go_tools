package handler

import (
	"fmt"
	"net/http"
	TM "odin_tool/models/twod/master"
)

type WorldController struct {
	Controller
}

func (c WorldController) List(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	world := TM.World{}
	worlds := world.LoadAll()

	c.data = make(map[string]interface{})
	c.data["worlds"] = worlds
	c.Render(w, r)
}
