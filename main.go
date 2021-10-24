package main

import (
	"Gin/routers"
	"Gin/utils"
	"fmt"
)

func main() {
	listen_port := fmt.Sprintf(":%v", utils.ServerConf.ListenPort)

	engin := routers.NewRouter()
	engin.Run(listen_port)
}
