package main

import (
	"net/http"
	"odin_tool/handler"
	"odin_tool/handler/admin"
	"odin_tool/libs"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		// handler = libs.Casbin(handler, route.Name, route.Method)
		handler = libs.SetCurrentUser(handler)
		handler = libs.SessionStart(handler)

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
	{"Error", "Get", "/error", handler.IndexController{}.Error},

	{"AuthLogin", "Get", "/auth/login", handler.AuthController{}.Login},
	{"AuthLogin", "Post", "/auth/login", handler.AuthController{}.DoLogin},
	{"AuthLogout", "Get", "/auth/logout", handler.AuthController{}.Logout},

	{"WorldList", "Get", "/world/list", handler.WorldController{}.List},
	{"AreaList", "Get", "/area/list", handler.AreaController{}.List},
	{"StageList", "Get", "/stage/list", handler.StageController{}.List},
	{"StageDetail", "Get", "/stage/{id:[0-9]+}", handler.StageController{}.Get},

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

	//mission
	{"MissionDailyList", "Get", "/mission/daily/list", handler.MissionController{}.DailyList},

	//monster
	{"MonsterList", "Get", "/monster/list", handler.MonsterController{}.List},

	//admin
	{"AdminUserList", "Get", "/admin/user/list", admin.UserController{}.List},
	{"AdminUserNew", "Get", "/admin/user/new", admin.UserController{}.New},
	{"AdminUserNew", "Post", "/admin/user/new", admin.UserController{}.Create},
	{"AdminUserRules", "Get", "/admin/user/{id:[0-9]+}/rules", admin.UserController{}.Rules},
	{"AdminUserRules", "Get", "/admin/user/{id:[0-9]+}/rules/add", admin.UserController{}.AddRules},
	{"AdminUserRules", "Get", "/admin/user/{id:[0-9]+}/rules/del", admin.UserController{}.DelRules},
}
