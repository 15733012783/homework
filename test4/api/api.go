package api

import (
	"github.com/15733012783/proto"
	"google.golang.org/grpc"
)

func Register(s grpc.ServiceRegistrar) {
	proto.RegisterGoodsServer(s, GoodsService{})
}
