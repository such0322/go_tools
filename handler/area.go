package handler

import (
	"fmt"
	"net/http"
	TM "odin_tool/models/twod/master"
	"strconv"
)

type AreaController struct {
	Controller
}

func (c AreaController) List(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	worldID, err := strconv.Atoi(r.FormValue("world_id"))
	if err != nil {
		panic("参数错误")
	}

	area := TM.Area{}
	areas := area.LoadByWorldID(worldID)

	c.data = make(map[string]interface{})
	c.data["areas"] = areas
	c.Render(w, r)
}
