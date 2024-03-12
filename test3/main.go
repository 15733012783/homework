package main

import (
	"2108-a-homework/test3/api"
	"flag"
	"github.com/15733012783/grpc/grpc"
	"github.com/15733012783/mysql/model"
	grpc2 "google.golang.org/grpc"
)

var (
	prot = flag.Int("port", 8081, "The server port")
)

func main() {
	flag.Parse()
	model.InitMysql()
	grpc.RegisterGRPC(*prot, func(s *grpc2.Server) {
		api.Register(s)
	})
}
