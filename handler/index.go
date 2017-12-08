package handler

import (
	"net/http"
)

type IndexController struct {
	Controller
}

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	c.tpl = "/index"
	c.Render(w, r)

}
