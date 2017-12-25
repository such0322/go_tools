package libs

import (
	"net/http"
	"odin_tool/libs/session"
	_ "odin_tool/libs/session/provider"
)

var globalSessions *session.Manager
var Session session.Session

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func SessionStart(inner http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Session = globalSessions.SessionStart(w, r)
		inner.ServeHTTP(w, r)
	})
}
