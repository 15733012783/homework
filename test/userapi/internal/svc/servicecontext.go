package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"test/userapi/internal/config"
	"test/userrpc/userrpcclient"
)

type ServiceContext struct {
	Config  config.Config
	Userrpc userrpcclient.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Userrpc: userrpcclient.NewUserrpc(zrpc.MustNewClient(c.Userrpc)),
	}
}
