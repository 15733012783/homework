package api

import (
	"2108-a-homework/test4/service"
	"context"
	"fmt"
	"github.com/15733012783/proto"
	"google.golang.org/grpc/status"
)

type GoodsService struct {
	proto.UnimplementedGoodsServer
}

func (GoodsService) CreateGoods(ctx context.Context, in *proto.CreateGoodsRequest) (*proto.CreateGoodsResponse, error) {
	if in == nil {
		return nil, status.Error(0, "值不能为空")
	}
	//if in.Info.ID == 0 {
	//	return nil, status.Error(0, "ID不能为0")
	//}
	goods, err := service.CreateGoods(in.Info)
	if err != nil {
		return nil, status.Error(0, err.Error())
	}
	fmt.Println("********************************")
	return &proto.CreateGoodsResponse{
		Info: goods,
	}, nil
}

func (GoodsService) DeleteGoods(ctx context.Context, in *proto.DeleteGoodsRequest) (*proto.DeleteGoodsResponse, error) {
	if in.ID == 0 {
		return nil, status.Error(0, "ID不能为0")
	}
	err := service.DeleteGoods(int(in.ID))
	if err != nil {
		return nil, status.Error(0, err.Error())
	}
	return &proto.DeleteGoodsResponse{}, nil
}

func (GoodsService) UploadGoods(ctx context.Context, in *proto.UploadGoodsRequest) (*proto.UploadGoodsResponse, error) {
	if in == nil {
		return nil, status.Error(0, "值不能为空")
	}
	if in.Info.ID == 0 {
		return nil, status.Error(0, "ID不能为0")
	}
	goods, err := service.UploadGoods(in.Info)
	if err != nil {
		return nil, status.Error(0, err.Error())
	}
	return &proto.UploadGoodsResponse{
		Info: goods,
	}, nil
}

func (GoodsService) WhereGoods(ctx context.Context, in *proto.WhereGoodsRequest) (*proto.WhereGoodsResponse, error) {
	if in.GoodsName == "" {
		return nil, status.Error(0, "搜索值不能为空")
	}
	goods, err := service.WhereGoods(in.GoodsName)
	if err != nil {
		return nil, status.Error(0, err.Error())
	}
	return &proto.WhereGoodsResponse{
		Info: goods,
	}, nil
}

func (GoodsService) UploadFile(ctx context.Context, in *proto.UploadFileRequest) (*proto.UploadFileResponse, error) {
	if in.ID == 0 {
		return nil, status.Error(0, "ID不能为空")
	}
	goods, err := service.UploadFile(int(in.ID), in.GoodsName)
	if err != nil {
		return nil, status.Error(0, err.Error())
	}
	return &proto.UploadFileResponse{
		Info: goods,
	}, nil
}
