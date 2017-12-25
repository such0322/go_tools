package handler

import (
	"net/http"
)

type FeatureController struct {
	Controller
}

func (c FeatureController) Search(w http.ResponseWriter, r *http.Request) {
	c.Tpl = "feature/search"
	c.Render(w, r)
}

func (c FeatureController) List(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// model := master.Features{}
	// features, _ := model.GetPage(page)
	//
	// c.data = make(map[string]interface{})
	// c.data["features"] = features
	// c.tpl = "feature/list"
	// c.Render(w, r)

}

func (c FeatureController) Get(w http.ResponseWriter, r *http.Request) {

}

func (c FeatureController) Update(w http.ResponseWriter, r *http.Request) {

}
