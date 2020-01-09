package passage_master

/*
Http handler registration module
*/

import "net/http"

func httpHelloWorldHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resp := helloWorldHandler()
		w.Write(resp)
	}
	return http.HandlerFunc(fn)
}

func httpHeartbeatHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resp := heartbeatHandler()
		w.Write(resp)
	}
	return http.HandlerFunc(fn)
}

func DispatchHttpApiHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/helloworld", httpHelloWorldHandler())
	mux.Handle("/heartbeat", httpHeartbeatHandler())
	return mux
}