package server

import (
	"google.golang.org/grpc"
	_ "github.com/golang/protobuf/proto"
)

type serverGrpcClient struct {}

func (s serverGrpcClient) HeartbeatMaster() {
	/*
		Renews server heartbeat with the master
	*/
}

func (s serverGrpcClient) HeartbeatClient() {
	/*
		Ensures that the client is up-to-date via heart beating its proxy
	*/
}

func (s serverGrpcClient) AddClient() {
	/*
		Registers a new client proxy connected to the server.
		In the process, we maintain heartbeats with the client via his queue
	*/
}

func (s serverGrpcClient) RemoveClient() {
	/*
		De-registers a client upon request.
	*/
}

func (s serverGrpcClient) Dispatch() {
	/*
		Dispatches a message to another app registered with other passage servers.
		We assume that all apps registered under the same name are the same.
		Payload is first serialized to a byte array before being wrapped in a server specific message structure and transmitted to destination
	*/
}

func (s serverGrpcClient) ListAddressesByName() {
	/*
		Lists the availability of an app.
		Target app is identified by its application name
	*/
}

func DispatchGrpcApiHandler() *grpc.Server {
	grpcServer := grpc.NewServer()
	return grpcServer
}