package server

import (
	"github.com/golang/protobuf/ptypes/empty"
	"go-passage/pkg/passage-master/dao"
	"go-passage/pkg/passage-master/handlers"
	"google.golang.org/grpc"
	"context"
	pb "go-passage/api/protobuf-spec/bin/master"
)

type grpcClient struct {
	Handler handlers.AccessHandler
}

func (g grpcClient) HelloWorld(ctx context.Context, empty *empty.Empty) (*pb.HelloWorldResponse, error) {
	resp := "Hello world"
	return &pb.HelloWorldResponse{Resp:resp}, nil
}

func (g grpcClient) Heartbeat(ctx context.Context, empty *empty.Empty) (*pb.HeartbeatResponse, error) {
	resp := true
	return &pb.HeartbeatResponse{Resp:resp}, nil
}

/*
	Basic CRUD operations for Servers to register themselves, as well as the applications that are docked to them
*/

func (g grpcClient) AddServer(ctx context.Context, request *pb.AddServerRequest) (*pb.AddServerResponse, error) {
	/*
		Registers a new Passage server. There can only be one master but multiple servers.

		Params
			- request: 	Add server request
		Returns
			- pb.AddServerResponse: proto add server response
			- error: Error when adding server
	*/
	err := g.Handler.RegisterServer(request.ExposedUrl, int(request.ExposedPort)); if err != nil {return nil, err}

	srv, err := g.Handler.SearchServer(request.ExposedUrl, int(request.ExposedPort)); if err != nil {
		return nil, err
	} else {
		return &pb.AddServerResponse{
			Resp:     true,
			ServerID: uint32(srv.ID),
		},nil
	}
}

func (g grpcClient) RemoveServer(ctx context.Context, request *pb.RemoveServerRequest) (*pb.RemoveServerResponse, error) {
	/*
		Removes an existing Passage server

		Params
			- request: 	Remove server request
		Returns
			- pb.RemoveServerResponse: proto remove server response
			- error: Error when adding server
	*/
	err := g.Handler.DeleteServer(int(request.ServerID)); if err != nil {return nil, err}
	return &pb.RemoveServerResponse{Resp: true},nil
}

func (g grpcClient) AddApplication(ctx context.Context, request *pb.AddAppRequest) (*pb.AddAppResponse, error) {
	/*
		Adds an existing application registered already to a server. Each application is identified by its name.
		There may be multiple applications sharing the same name but only one (serverID, appName) pair.

		Params
			- request: 	Add application request
		Returns
			- pb.AddAppResponse: proto add application response
			- error: Error when adding server
	*/
	err := g.Handler.RegisterApp(int(request.ServerID), request.AppName); if err != nil {return nil, err}
	return &pb.AddAppResponse{Resp: true}, nil
}

func (g grpcClient) RemoveApplication(ctx context.Context, request *pb.RemoveAppRequest) (*pb.RemoveAppResponse, error) {
	/*
		Removes an application registered already to a server, identified by unique (serverID, appName) pair

		Params
			- request: 	Remove application request
		Returns
			- pb.AddAppResponse: proto add application response
			- error: Error when adding server
	*/
	return nil,nil
}

func (g grpcClient) CheckServeIsAlive(ctx context.Context, empty *empty.Empty) (*pb.HeartbeatResponse, error) {
	/*
		Asserts liveliness check by pinging server

		Params
			- nil
		Returns
			- pb.HeartbeatResponse: proto heartbeat response from pinged server
			- error: Error when attempting to ping server
	*/
	return nil,nil
}

func (g grpcClient) GetAppAddress(ctx context.Context, request *pb.GetAppRequest) (*pb.GetAppResponse, error) {
	/*
		Returns address of the application present within each server, if the application does exist.
		Else, we return an empty response

		Params
			- request: Get Application address request
		Returns
			- pb.GetAppResponse: If application exists, return the message type GetAppResponse, else return nothing
			- error: Error when attempting to retrieve application
	*/
	return nil,nil
}

func DispatchGrpcApiHandler() *grpc.Server {
	grpcServer := grpc.NewServer()
	dataObj := dao.MysqlDAO{}
	pb.RegisterMasterServicesServer(grpcServer, &grpcClient{Handler: handlers.AccessHandler{ModelDAO: dataObj}})
	return grpcServer
}