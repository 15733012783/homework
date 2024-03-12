package api

import (
	"2108-a-homework/test3/service"
	"context"
	"github.com/15733012783/proto"

	"google.golang.org/grpc/status"
)

type GoodsService struct {
	proto.UnimplementedGoodsServer
}

func (GoodsService) CreateGoods(ctx context.Context, in *proto.CreateGoodsRequest) (*proto.CreateGoodsResponse, error) {
	if in == nil {
		return nil, status.Error(0, "不能为空")
	}
	if in.Info.ID == 0 {
		return nil, status.Error(0, "id不能为空并且大于0")
	}
	goods, err := service.CreateGoods(in.Info)
	if err != nil {
		return nil, status.Error(200, err.Error())
	}
	return &proto.CreateGoodsResponse{
		Info: goods,
	}, err
}

func (GoodsService) DeleteGoods(ctx context.Context, in *proto.DeleteGoodsRequest) (*proto.DeleteGoodsResponse, error) {
	if in == nil {
		return nil, status.Error(0, "不能为空")
	}
	if in.ID == 0 {
		return nil, status.Error(0, "id不能为空并且大于0")
	}
	err := service.DeleteGoods(int(in.ID))
	if err != nil {
		return nil, status.Error(200, err.Error())
	}
	return &proto.DeleteGoodsResponse{}, err
}

func (GoodsService) UploadGoods(ctx context.Context, in *proto.UploadGoodsRequest) (*proto.UploadGoodsResponse, error) {
	if in == nil {
		return nil, status.Error(0, "不能为空")
	}
	if in.Info.ID == 0 {
		return nil, status.Error(0, "id不能为空并且大于0")
	}
	goods, err := service.UploadGoods(in.Info)
	if err != nil {
		return nil, status.Error(200, err.Error())
	}
	return &proto.UploadGoodsResponse{
		Info: goods,
	}, err
}

func (GoodsService) WhereGoods(ctx context.Context, in *proto.WhereGoodsRequest) (*proto.WhereGoodsResponse, error) {
	if in == nil {
		return nil, status.Error(0, "不能为空")
	}
	if in.GoodsName == "" {
		return nil, status.Error(0, "用户名不能为空并且大于0")
	}
	goods, err := service.WhereGoods(in.GoodsName)
	if err != nil {
		return nil, status.Error(200, err.Error())
	}
	return &proto.WhereGoodsResponse{
		Info: goods,
	}, err
}
