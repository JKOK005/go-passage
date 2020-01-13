package main

import (
	"flag"
	"go-passage/pkg/passage-master/server"
	"sync"
)

func main () {
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		/*
			Start http server
		*/
		srv := server.MasterServer{
			HttpServerUrl:  "localhost",
			HttpServerPort: 9000,
		}
		_, err := srv.StartHttpServer(); if err != nil {panic(err)}
	}()

	go func(){
		/*
			Start gRPC server
		*/
		srv := server.MasterServer{
			GrpcServerUrl:  "localhost",
			GrpcServerPort: 9001,
		}
		_, err := srv.StartGrpcServer(); if err != nil {panic(err)}
	}()

	wg.Wait()
}
