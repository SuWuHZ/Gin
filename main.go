package main

import (
	"Gin/routers"
	"Gin/utils"
	"fmt"
)

func main() {

	serverConfig := utils.ServerConf
	listen_port := fmt.Sprintf(":%v", serverConfig.Server.Port)
	engin := routers.NewRouter()
	engin.Run(listen_port)
}
