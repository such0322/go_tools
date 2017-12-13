package handler

import (
	"net/http"
)

type AuthController struct {
	Controller
}

func (c Controller) Login(w http.ResponseWriter, r *http.Request) {

	c.Render(w, r)
}
