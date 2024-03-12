package test2

import (
	"2108-a-homework/test2/logic"
	"2108-a-homework/test2/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", "0.0.0.0:8080")
	server := grpc.NewServer()
	var s logic.S
	proto.RegisterStreamGreeterServer(server, &s)
	err := server.Serve(lis)
	if err != nil {
		return
	}
}
