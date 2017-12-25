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
	Data map[string]interface{}
	Tpl  string
}

func test(a interface{}) string {
	return "this is test"
}

var funcMap = template.FuncMap{
	"test":       test,
	"jsonDecode": libs.JsonDecode,
	"inIntSlice": libs.InIntSlice,
}

func (c *Controller) Render(w http.ResponseWriter, r *http.Request) {
	c.Data["currentUser"] = libs.CurrentUser
	if c.Tpl == "" {
		c.Tpl = r.URL.Path
	}
	tpl := strings.Split(c.Tpl, "/")
	file := bytes.Buffer{}
	file.WriteString("templates")
	file.WriteString(c.Tpl)
	file.WriteString(".html")

	tmpl := template.Must(template.New(tpl[len(tpl)-1]).Funcs(funcMap).ParseFiles(file.String(), "templates/header.html", "templates/footer.html"))
	if err := tmpl.ExecuteTemplate(w, c.Tpl, c.Data); err != nil {
		log.Fatal(err)
	}

}
