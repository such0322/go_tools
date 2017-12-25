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
		fmt.Println("参数错误")
	}
	area := TM.Area{}
	areas := TM.Areas{}
	if worldID != 0 {
		areas = area.GetByWorldID(worldID)
	} else {
		areas = area.GetAll()
	}

	c.Data = make(map[string]interface{})
	c.Data["areas"] = areas
	c.Render(w, r)
}
