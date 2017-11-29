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

func mapValue(m map[interface{}]interface{}, key interface{}) interface{} {
	return m[key]
}

var funcMap = template.FuncMap{
	"mapValue": mapValue,
}

func (c *Controller) Render(w http.ResponseWriter, r *http.Request) {
	if c.tpl == "" {
		c.tpl = r.URL.Path
	}
	tpl := strings.Split(c.tpl, "/")
	file := bytes.Buffer{}
	file.WriteString("templates")
	file.WriteString(c.tpl)
	file.WriteString(".html")

	tmpl := template.Must(template.New(tpl[len(tpl)-1]).ParseFiles(file.String(), "templates/header.html", "templates/footer.html"))
	tmpl.Funcs(funcMap)
	if err := tmpl.ExecuteTemplate(w, c.tpl, c.data); err != nil {
		log.Fatal(err)
	}

}
