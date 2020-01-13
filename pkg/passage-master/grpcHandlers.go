package passage_master

import (
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"context"
	pb "go-passage/api/protobuf-spec/bin/master"
)

type grpcClient struct {}

func (g *grpcClient) HelloWorld(ctx context.Context, empty *empty.Empty) (*pb.HelloWorldResponse, error) {
	resp := "Hello world"
	return &pb.HelloWorldResponse{Resp:resp}, nil
}

func (g *grpcClient) Heartbeat(ctx context.Context, empty *empty.Empty) (*pb.HeartbeatResponse, error) {
	resp := true
	return &pb.HeartbeatResponse{Resp:resp}, nil
}

/*
	Basic CRUD operations for Servers to register themselves, as well as the applications that are docked to them
*/

func (g *grpcClient) AddServer(ctx context.Context, request *pb.AddServerRequest) (*pb.AddServerResponse, error) {
	/*
		Registers a new Passage server. There can only be one master but multiple servers.

		Params
			- request: 	Add server request
		Returns
			- pb.AddServerResponse: proto add server response
			- error: Error when adding server
	*/
	return nil,nil
}

func (g *grpcClient) RemoveServer(ctx context.Context, request *pb.RemoveServerRequest) (*pb.RemoveServerResponse, error) {
	/*
		Removes an existing Passage server

		Params
			- request: 	Remove server request
		Returns
			- pb.RemoveServerResponse: proto remove server response
			- error: Error when adding server
	*/
	return nil,nil
}

func (g *grpcClient) AddApplication(ctx context.Context, request *pb.AddAppRequest) (*pb.AddAppResponse, error) {
	/*
		Adds an existing application registered already to a server. Each application is identified by its name.
		There may be multiple applications sharing the same name but only one (serverID, appName) pair.

		Params
			- request: 	Add application request
		Returns
			- pb.AddAppResponse: proto add application response
			- error: Error when adding server
	*/
	return nil,nil
}

func (g *grpcClient) RemoveApplication(ctx context.Context, request *pb.RemoveAppRequest) (*pb.RemoveAppResponse, error) {
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

func (g *grpcClient) CheckServeIsAlive(ctx context.Context, empty *empty.Empty) (*pb.HeartbeatResponse, error) {
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

func (g *grpcClient) GetAppAddress() {}

func DispatchGrpcApiHandler() *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterMasterServicesServer(grpcServer, &grpcClient{})
	return grpcServer
}