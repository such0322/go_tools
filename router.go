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
	{"Index", "Get", "/", handler.IndexController{}.Index},

	{"AuthLogin", "Get", "/auth/login", handler.AuthController{}.Login},

	{"StageList", "Get", "/stage/list", handler.StageController{}.List},
	{"StageDetail", "Get", "/stage/{id:[0-9]+}", handler.StageController{}.Get},
	// {"Login", "Get", "/login", handler.Login},
	// {"FeatureTool", "Get", "/feature/tool", handler.ToolsController{}.Tool},
	{"FeatureSearch", "Get", "/feature/search", handler.FeatureController{}.Search},
	{"FeatureList", "Get", "/feature/list", handler.FeatureController{}.List},
	{"FeatureDetail", "Get", "/feature/detail", handler.FeatureController{}.Get},
	{"FeatureUpdate", "Put", "/feature/update", handler.FeatureController{}.Update},

	//gift
	{"GiftList", "Get", "/gift/list", handler.GiftController{}.List},
	{"GiftNew", "Get", "/gift/new", handler.GiftController{}.NewGift},
	{"GiftNew", "Post", "/gift/new", handler.GiftController{}.Create},
	{"GiftRandomCode", "Get", "/gift/randomCode", handler.GiftController{}.RandomCode},
	{"GiftBounsAll", "Get", "/gift/getBounsAll", handler.GiftController{}.GetBounsAll},

	//monster
	{"MonsterList", "Get", "/monster/list", handler.MonsterController{}.List},
}
