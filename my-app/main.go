package main

import (
	"2108-a-homework/my-app/routers"
	"log"
)

func main() {
	err := routers.Router.Run(":8080")
	if err != nil {
		log.Println(err, "启动失败")
		return
	}
}
