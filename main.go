package main

import (
	"flag"
	_ "fmt"
	_ "go-passage/pkg/passage-master/schema"
	"go-passage/pkg/passage-master/server"
	"sync"
)

func main () {
	flag.Parse()
	//
	//
	//go func(){
	//	/*
	//		Start http server
	//	*/
	//	srv := server.MasterServer{
	//		HttpServerUrl:  "localhost",
	//		HttpServerPort: 9000,
	//	}
	//	_, err := srv.StartHttpServer(); if err != nil {panic(err)}
	//}()
	//

	// Trigger migration
	//schema.StartMigration()

	// Test CRUD for server model
	//dobj := dao.MysqlDAO{}
	//
	//newSrv := schema.ServerModel{
	//	Url:   "localhost2",
	//	Port:  9010,
	//}
	//dobj.CreateServer(&newSrv)

	//srv, _ := dobj.GetServer(map[string]string{"url":"localhost"})
	//dobj.DeleteServer(srv)

	//newApp := schema.AppModel{
	//	Name:   "testApp5",
	//	DeployedServerUid: srv.ID,
	//}
	//dobj.CreateApp(&newApp)
	//
	//dobj.GetApps(map[string]string{"name":"testApp5"})

	//db, _ := gorm.Open("mysql", "passage:passage@(localhost)/Passage_master?charset=utf8&parseTime=True&loc=Local")
	//var y []schema.AppModel
	//db.Model(srv).Related(y)
	//fmt.Printf("asd %v", y)

	// Test grpc
	// Start the client

	var wg sync.WaitGroup
	wg.Add(2)

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

	// Add server procedure
	//if conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 9001), grpc.WithInsecure()); err != nil {
	//	log.Println(err)
	//} else {
	//	defer conn.Close()
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000) * time.Millisecond)
	//	defer cancel()
	//
	//	client := pb.NewMasterServicesClient(conn)
	//
	//	resp, err := client.AddServer(ctx, &pb.AddServerRequest{
	//		ExposedUrl:           "localhost",
	//		ExposedPort:          8001,
	//	})
	//
	//	if ctx.Err() == context.DeadlineExceeded {
	//		// Request timed out. Report as timeout.
	//		log.Println("Request timed out: ", ctx.Err())
	//	}else {
	//		// Request succeeded
	//		if err != nil {
	//			glog.Error(err)
	//		} else {
	//			glog.Info("Request from client succeeded. Client ID: ", resp.ServerID)
	//		}
	//	}
	//}
	//
	//// Add App
	//if conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 9001), grpc.WithInsecure()); err != nil {
	//	log.Println(err)
	//} else {
	//	defer conn.Close()
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000) * time.Millisecond)
	//	defer cancel()
	//
	//	client := pb.NewMasterServicesClient(conn)
	//
	//	resp, err := client.AddApplication(ctx, &pb.AddAppRequest{
	//		ServerID: 11,
	//		AppName: "testApp1",
	//	})
	//
	//	if ctx.Err() == context.DeadlineExceeded {
	//		// Request timed out. Report as timeout.
	//		log.Println("Request timed out: ", ctx.Err())
	//	}else {
	//		// Request succeeded
	//		if err != nil {
	//			glog.Error(err)
	//		} else {
	//			glog.Info("Add app status: ", resp.Resp)
	//		}
	//	}
	//}

	// Validate app deployment
	//if conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 9001), grpc.WithInsecure()); err != nil {
	//	log.Println(err)
	//} else {
	//	defer conn.Close()
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000) * time.Millisecond)
	//	defer cancel()
	//
	//	client := pb.NewMasterServicesClient(conn)
	//
	//	resp, err := client.GetAppAddress(ctx, &pb.GetAppRequest{
	//		AppName:              "testApp1",
	//	})
	//
	//	if ctx.Err() == context.DeadlineExceeded {
	//		// Request timed out. Report as timeout.
	//		log.Println("Request timed out: ", ctx.Err())
	//	}else {
	//		// Request succeeded
	//		if err != nil {
	//			glog.Error(err)
	//		} else {
	//			for _, srv := range resp.SrvAddr{
	//				glog.Info(fmt.Printf("App deployed at %s:%d \n", srv.ServerUrl, srv.ServerPort))
	//			}
	//		}
	//	}
	//}

	// Remove application procedure
	//if conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 9001), grpc.WithInsecure()); err != nil {
	//	log.Println(err)
	//} else {
	//	defer conn.Close()
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000) * time.Millisecond)
	//	defer cancel()
	//
	//	client := pb.NewMasterServicesClient(conn)
	//
	//	resp, err := client.RemoveApplication(ctx, &pb.RemoveAppRequest{
	//		ServerID: 11,
	//		AppName: "testApp1",
	//	})
	//
	//	if ctx.Err() == context.DeadlineExceeded {
	//		// Request timed out. Report as timeout.
	//		log.Println("Request timed out: ", ctx.Err())
	//	}else {
	//		// Request succeeded
	//		if err != nil {
	//			glog.Error(err)
	//		} else {
	//			glog.Info(fmt.Printf("App removed status: %b \n", resp.Resp))
	//		}
	//	}
	//}

	// Remove server procedure
	//if conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 9001), grpc.WithInsecure()); err != nil {
	//	log.Println(err)
	//} else {
	//	defer conn.Close()
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000) * time.Millisecond)
	//	defer cancel()
	//
	//	client := pb.NewMasterServicesClient(conn)
	//
	//	resp, err := client.RemoveServer(ctx, &pb.RemoveServerRequest{
	//		ServerID: 3,
	//	})
	//
	//	if ctx.Err() == context.DeadlineExceeded {
	//		// Request timed out. Report as timeout.
	//		log.Println("Request timed out: ", ctx.Err())
	//	}else {
	//		// Request succeeded
	//		if err != nil {
	//			glog.Error(err)
	//		} else {
	//			glog.Info("Removed server with: ", resp.Resp)
	//		}
	//	}
	//}
	wg.Wait()
}
