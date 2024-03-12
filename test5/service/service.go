package service

import (
	"2108-a-homework/test5/client"
	"context"
	"fmt"
	"github.com/15733012783/proto"
)

func CreateGoods(ctx context.Context, in *proto.GoodsInfo) (*proto.GoodsInfo, error) {
	goods, err := client.CreateGoods(ctx, in)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, fmt.Errorf("商品添加失败")
	}
	return goods, nil
}

func DeleteGoods(ctx context.Context, id int64) error {
	goods, err := client.DeleteGoods(ctx, id)
	if err != nil {
		return err
	}
	if goods == nil {
		return fmt.Errorf("商品删除失败")
	}
	return nil
}

func UpdateGoods(ctx context.Context, in *proto.GoodsInfo) (interface{}, error) {
	goods, err := client.UploadGoods(ctx, in)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, fmt.Errorf("商品删除失败")
	}
	return goods, nil
}

func WhereGoods(ctx context.Context, GoodsName string) (interface{}, error) {
	goods, err := client.WhereGoods(ctx, GoodsName)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, fmt.Errorf("商品删除失败")
	}
	return goods, nil
}

func UploadFile(ctx context.Context, id int64, fileName string) (*proto.GoodsInfo, error) {
	goods, err := client.UploadFile(ctx, id, fileName)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, fmt.Errorf("文件上传失败")
	}
	return goods, nil
}
