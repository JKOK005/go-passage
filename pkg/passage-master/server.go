package passage_master

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
)

type MasterServer struct
{
	HttpServerUrl 	string
	HttpServerPort 	int
}

func (m *MasterServer) StartHttpServer () (int, error) {
	/*
	Begins http server setup. This is used to service application requests via http api calls.

	Params
		- nil
	Returns
		- int: 		Http status code
		- error: 	Error when initializing server
	*/
	glog.Infof(fmt.Sprintf("Initializing master http server at %s:%d", m.HttpServerUrl, m.HttpServerPort))
	mux := DispatchApiHandler()
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", m.HttpServerUrl, m.HttpServerPort), mux); if err != nil {
		return http.StatusNotFound, err
	}
	return http.StatusAccepted, nil
}