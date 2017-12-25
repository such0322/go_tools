package libs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"odin_tool/models/tool"
)

var CurrentUser tool.User

func SetCurrentUser(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonUser := Session.Get("currentUser")
		if jsonUser != nil {
			err := json.Unmarshal([]byte(jsonUser.(string)), &CurrentUser)
			if err != nil {
				fmt.Println(err)
			}
		}
		inner.ServeHTTP(w, r)
	})
}
