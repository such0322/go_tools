package handler

import (
	"net/http"
	"odin_tools/models"
)

func Tool(w http.ResponseWriter, r *http.Request) {
	var dc models.DeviceCredential
	dc.GetByGuid()
	// var t models.Ttt
	// t.GetAll()

}
