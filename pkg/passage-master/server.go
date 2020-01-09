package passage_master

import (
	"fmt"
	"github.com/golang/glog"
	"net"
	"net/http"
)

type MasterServer struct
{
	HttpServerUrl 	string
	HttpServerPort 	int
	GrpcServerUrl 	string
	GrpcServerPort 	int
}

func (m *MasterServer) StartHttpServer () (bool, error) {
	/*
	Begins http server setup. This is used to service application requests via http api calls.

	Params
		- nil
	Returns
		- bool: 	Is success
		- error: 	Error when initializing server
	*/
	glog.Infof(fmt.Sprintf("Initializing master http server at %s:%d", m.HttpServerUrl, m.HttpServerPort))
	mux := DispatchHttpApiHandler()
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", m.HttpServerUrl, m.HttpServerPort), mux); if err != nil {
		glog.Fatal(err)
		return false, err
	}
	return true, nil
}

func (m *MasterServer) StartGrpcServer() (bool, error) {
	/*
		Begins gRPC server setup. This is used to service application requests via Passage servers.

		Params
			- nil
		Returns
			- bool: 	Is success
			- error: 	Error when initializing server
	*/
	glog.Infof(fmt.Sprintf("Initializing master gRPC server at %s:%d", m.GrpcServerUrl, m.GrpcServerPort))
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", m.GrpcServerUrl, m.GrpcServerPort)); if err != nil {
		glog.Fatal(err)
		return false, err
	}
	grpcServer := DispatchGrpcApiHandler()
	grpcServer.Serve(lis)
	return true, nil
}