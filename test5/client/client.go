package client

import (
	"context"
	"fmt"
	"github.com/15733012783/mysql/client"
	"github.com/15733012783/mysql/nacos"
	"github.com/15733012783/proto"
	"log"
)

type handler func(ctx context.Context, cli proto.GoodsClient) (interface{}, error)

func withClient(ctx context.Context, handler handler) (interface{}, error) {
	sprintf := fmt.Sprintf("%v:%v", nacos.GoodsT.Consul.Host, 8500)
	fmt.Println(sprintf, "*****")
	conn, err := client.Client(nacos.GoodsT.Consul.Name, sprintf)
	fmt.Println(conn, "conn")
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	goodsCli := proto.NewGoodsClient(conn)
	res, err := handler(ctx, goodsCli)
	if err != nil {
		return nil, err
	}
	conn.Close()
	return res, nil
}

//func withClients(ctx context.Context, handler handler) (interface{}, error) {
//	dial, err := grpc.Dial(fmt.Sprintf("%v:%v", "10.2.171.70", "8881"), grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		return nil, err
//	}
//	config := proto.NewGoodsClient(dial)
//	res, err := handler(ctx, config)
//	if err != nil {
//		return nil, err
//	}
//	dial.Close()
//	return res, nil
//}

func CreateGoods(ctx context.Context, in *proto.GoodsInfo) (*proto.GoodsInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli proto.GoodsClient) (interface{}, error) {
		goods, err := cli.CreateGoods(ctx, &proto.CreateGoodsRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return goods.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*proto.GoodsInfo), nil
}

func DeleteGoods(ctx context.Context, id int64) (interface{}, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli proto.GoodsClient) (interface{}, error) {
		goods, err := cli.DeleteGoods(ctx, &proto.DeleteGoodsRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return goods, nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func UploadGoods(ctx context.Context, in *proto.GoodsInfo) (interface{}, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli proto.GoodsClient) (interface{}, error) {
		goods, err := cli.UploadGoods(ctx, &proto.UploadGoodsRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return goods, nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func WhereGoods(ctx context.Context, GoodsName string) (interface{}, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli proto.GoodsClient) (interface{}, error) {
		goods, err := cli.WhereGoods(ctx, &proto.WhereGoodsRequest{
			GoodsName: GoodsName,
		})
		if err != nil {
			return nil, err
		}
		return goods, nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func UploadFile(ctx context.Context, id int64, filename string) (*proto.GoodsInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli proto.GoodsClient) (interface{}, error) {
		goods, err := cli.UploadFile(ctx, &proto.UploadFileRequest{
			ID:        id,
			GoodsName: filename,
		})
		if err != nil {
			return nil, err
		}
		return goods, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*proto.GoodsInfo), nil
}
