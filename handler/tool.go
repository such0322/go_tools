package handler

import (
	"fmt"
	"net/http"
	"odin_tools/models"
)

func Tool(w http.ResponseWriter, r *http.Request) {
	a := models.Artifact{}
	a.GetById(10002)
	fmt.Println(a)
}
