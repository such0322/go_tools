package handler

import (
	"fmt"
	"net/http"
)

type IndexController struct {
	Controller
}

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {

	c.Tpl = "/index"
	c.Data = make(map[string]interface{})
	c.Render(w, r)

}

func (c Controller) Error(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Printf("%#v\n", header)
	// c.tpl = "/error"
	// c.Render(w, r)
}
