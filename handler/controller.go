package handler

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"odin_tool/libs"
	"strings"
)

type Controller struct {
	data map[string]interface{}
	tpl  string
}

func test(a interface{}) string {
	return "this is test"
}

var funcMap = template.FuncMap{
	"test":       test,
	"jsonDecode": libs.JsonDecode,
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

	tmpl := template.Must(template.New(tpl[len(tpl)-1]).Funcs(funcMap).ParseFiles(file.String(), "templates/header.html", "templates/footer.html"))
	if err := tmpl.ExecuteTemplate(w, c.tpl, c.data); err != nil {
		log.Fatal(err)
	}

}
