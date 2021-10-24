package main

import (
	"fmt"
	"gin/routers"
)

func main() {
	fmt.Println("Hello SuWu!")
	engin := routers.NewRouter()
	engin.Run(":8001")
}
