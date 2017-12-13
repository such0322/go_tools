package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"odin_tool/models"
)

type ToolsController struct{}

func (c ToolsController) Tool(w http.ResponseWriter, r *http.Request) {
	as := models.Artifacts{}.GetAll()
	var data struct {
		As []models.Artifact
	}
	data.As = as

	tmpl := template.Must(template.New("tool").ParseFiles("templates/tool.html", "templates/header.html", "templates/footer.html"))
	fmt.Println(tmpl.Name())
	tmpl.ExecuteTemplate(w, "tool", `<b>World</b>`)

}
