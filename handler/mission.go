package handler

import "net/http"

type MissionController struct {
	Controller
}

func (c MissionController) DailyList(w http.ResponseWriter, r *http.Request) {

	c.data = make(map[string]interface{})
	c.Render(w, r)
}
