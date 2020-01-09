package passage_master

import (
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"context"
	pb "go-passage/api/protobuf-spec/bin/master"
)

type grpcClient struct {}

func (g *grpcClient) HelloWorld(ctx context.Context, empty *empty.Empty) (*pb.HelloWorldResponse, error) {
	resp := helloWorldHandler()
	return &pb.HelloWorldResponse{Resp:resp}, nil
}

func (g *grpcClient) Heartbeat(ctx context.Context, empty *empty.Empty) (*pb.HeartbeatResponse, error) {
	resp := heartbeatHandler()
	return &pb.HeartbeatResponse{Resp:resp}, nil
}

func DispatchGrpcApiHandler() *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterMasterServicesServer(grpcServer, &grpcClient{})
	return grpcServer
}