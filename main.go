package main

import (
	pm "go-passage/pkg/passage-master"
)

func main () {
	master := pm.MasterServer{
		HttpServerUrl:  "localhost",
		HttpServerPort: 9000,
	}
	master.StartHttpServer()
}
