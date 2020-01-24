package main

import (
	_ "fmt"
	"go-passage/pkg/passage-master/dao"
	_ "go-passage/pkg/passage-master/schema"
)

func main () {
	//flag.Parse()
	//
	//var wg sync.WaitGroup
	//wg.Add(2)
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
	//go func(){
	//	/*
	//		Start gRPC server
	//	*/
	//	srv := server.MasterServer{
	//		GrpcServerUrl:  "localhost",
	//		GrpcServerPort: 9001,
	//	}
	//	_, err := srv.StartGrpcServer(); if err != nil {panic(err)}
	//}()
	//
	//wg.Wait()

	// Trigger migration
	//schema.StartMigration()

	// Test CRUD for server model
	dobj := dao.MysqlDAO{}
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
	dobj.GetApps(map[string]string{"name":"testApp5"})

	//db, _ := gorm.Open("mysql", "passage:passage@(localhost)/Passage_master?charset=utf8&parseTime=True&loc=Local")
	//var y []schema.AppModel
	//db.Model(srv).Related(y)
	//fmt.Printf("asd %v", y)

}
