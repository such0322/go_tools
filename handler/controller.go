package handler

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Controller struct {
	data map[string]interface{}
	tpl  string
}

func (c *Controller) Render(w http.ResponseWriter, r *http.Request) {
	tpl := strings.Split(c.tpl, "/")
	file := bytes.Buffer{}
	file.WriteString("templates/")
	file.WriteString(c.tpl)
	file.WriteString(".html")
	tmpl := template.Must(template.New(tpl[1]).ParseFiles(file.String(), "templates/header.html", "templates/footer.html"))
	if err := tmpl.ExecuteTemplate(w, c.tpl, c.data); err != nil {
		log.Fatal(err)
	}
}
