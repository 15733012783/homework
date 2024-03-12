package main

import (
	"2108-a-homework/test4/api"
	"fmt"
	"github.com/15733012783/grpc/grpc"
	"github.com/15733012783/mysql/config"
	"github.com/15733012783/mysql/consul"
	"github.com/15733012783/mysql/model"
	"github.com/15733012783/mysql/nacos"
	grpc2 "google.golang.org/grpc"
	"strconv"
)

func main() {
	viper, err := config.InitViper("goods", "./config", "Nacos")
	if err != nil {
		return
	}
	atoi, err := strconv.Atoi(viper["port"])
	if err != nil {
		return
	}
	fmt.Println(viper)
	nacos.NaCosConfig(viper["group"], viper["dataid"], atoi)
	fmt.Println("", nacos.GoodsT)
	model.InItMysql()
	if err != nil {
		return
	}
	//port, err := consul.GetFreePort()
	//if err != nil {
	//	return
	//}
	ip := consul.GetIp()
	consul.SonSul(nacos.GoodsT.Grpc.Address, ip[0], nacos.GoodsT.Grpc.Port, nacos.GoodsT.Consul.Name)
	grpc.RegisterGRPC(ip[0], nacos.GoodsT.Grpc.Port, func(s *grpc2.Server) {
		api.Register(s)
	})
}
