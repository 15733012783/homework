package api

import (
	"github.com/15733012783/proto"
	"google.golang.org/grpc"
)

func Register(r grpc.ServiceRegistrar) {
	proto.RegisterGoodsServer(r, GoodsService{})
}
