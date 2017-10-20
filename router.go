package main

import (
	"net/http"
	"odin_tools/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		handler.Index,
	},
	Route{
		"Login",
		"Get",
		"/login",
		handler.Login,
	},
	Route{
		"FeatureTool",
		"Get",
		"/feature/tool",
		handler.Tool,
	},
}
