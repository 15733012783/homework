package main

import (
	"2108-a-homework/test5/api"
	"github.com/15733012783/mysql/nacos"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	api.Register(r)
	nacos.NaCosConfig("DEFAULT_GROUP", "1", 8848)
	err := r.Run(":8883")
	if err != nil {
		log.Println(err)
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
