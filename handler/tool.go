package handler

import (
	"html/template"
	"log"
	"net/http"
	"odin_tools/models"
)

func Tool(w http.ResponseWriter, r *http.Request) {
	as := models.Artifacts{}.GetAll()
	var data struct {
		As []models.Artifact
	}
	data.As = as
	temp := template.Must(template.New("tool.html").ParseFiles("templates/tool.html"))
	if err := temp.Execute(w, data); err != nil {
		log.Fatal(err)
	}

}
