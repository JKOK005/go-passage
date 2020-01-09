package passage_master

import (
	"google.golang.org/grpc"
)

func DispatchGrpcApiHandler() *grpc.Server {
	grpcServer := grpc.NewServer()
	return grpcServer
}