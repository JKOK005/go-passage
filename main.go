package main

import (
	"flag"
	pm "go-passage/pkg/passage-master"
)

func main () {
	flag.Parse()
	master := pm.MasterServer{
		HttpServerUrl:  "localhost",
		HttpServerPort: 9000,
	}
	master.StartHttpServer()
}
