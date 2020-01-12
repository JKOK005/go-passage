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

func (g *grpcClient) AddPassageServer() {}
func (g *grpcClient) RemovePassageServer() {}
func (g *grpcClient) AddApplication() {}
func (g *grpcClient) RemoveApplication() {}
func (g *grpcClient) CheckServeIsAlive() {}
func (g *grpcClient) GetAppAddress() {}

func DispatchGrpcApiHandler() *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterMasterServicesServer(grpcServer, &grpcClient{})
	return grpcServer
}