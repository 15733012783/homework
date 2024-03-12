package service

import (
	"github.com/15733012783/mysql/model"
	"github.com/15733012783/proto"
	"gorm.io/gorm"
)

func CreateGoods(in *proto.GoodsInfo) (*proto.GoodsInfo, error) {
	Mod := model.NewGoods()
	info, err := Mod.Create(GoodsMysql(in))
	if err != nil {
		return nil, err
	}
	return MsgMysql(info)
}

func DeleteGoods(id int) error {
	Mod := model.NewGoods()
	err := Mod.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func UploadGoods(in *proto.GoodsInfo) (*proto.GoodsInfo, error) {
	Mod := model.NewGoods()
	info, err := Mod.Upload(GoodsMysql(in))
	if err != nil {
		return nil, err
	}
	return MsgMysql(info)
}

func WhereGoods(GoodsName string) (*proto.GoodsInfo, error) {
	Mod := model.NewGoods()
	info, err := Mod.Get(GoodsName)
	if err != nil {
		return nil, err
	}
	return MsgMysql(info)
}

func UploadFile(id int, GoodsName string) (*proto.GoodsInfo, error) {
	Mod := model.NewGoods()
	info, err := Mod.UploadFile(id, GoodsName)
	if err != nil {
		return nil, err
	}
	return MsgMysql(info)
}

func GoodsMysql(in *proto.GoodsInfo) *model.Goods {
	return &model.Goods{
		Model: gorm.Model{
			ID: uint(in.ID),
		},
		GoodsName:  in.GoodsName,
		GoodsPrice: float64(in.GoodsPrice),
		GoodsNum:   int(in.GoodsNum),
	}
}

func MsgMysql(in *model.Goods) (*proto.GoodsInfo, error) {
	return &proto.GoodsInfo{
		ID:         int64(in.ID),
		GoodsName:  in.GoodsName,
		GoodsPrice: float32(in.GoodsPrice),
		GoodsNum:   int64(in.GoodsNum),
	}, nil
}
