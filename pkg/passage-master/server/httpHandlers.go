package server

/*
Http handlers registration module
*/

import (
	"net/http"
	"strconv"
)

func httpHelloWorldHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resp := "Hello world"
		w.Write([]byte(resp))
	}
	return http.HandlerFunc(fn)
}

func httpHeartbeatHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resp := true
		w.Write([]byte(strconv.FormatBool(resp)))
	}
	return http.HandlerFunc(fn)
}

func DispatchHttpApiHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/helloworld", httpHelloWorldHandler())
	mux.Handle("/heartbeat", httpHeartbeatHandler())
	return mux
}