package goods

import (
	"2108-a-homework/test5/service"
	"github.com/15733012783/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	var CreateGoodsReq struct {
		ID         int64
		GoodsName  string
		GoodsPrice float64
		GoodsNum   int64
	}
	err := c.ShouldBind(&CreateGoodsReq)
	if err != nil {
		Res(c, 10000, nil, err.Error())
		return
	}
	goods, err := service.CreateGoods(c, &proto.GoodsInfo{
		ID:         CreateGoodsReq.ID,
		GoodsName:  CreateGoodsReq.GoodsName,
		GoodsPrice: float32(CreateGoodsReq.GoodsPrice),
		GoodsNum:   CreateGoodsReq.GoodsNum,
	})
	if err != nil {
		Res(c, 10001, nil, err.Error())
		return
	}
	Res(c, http.StatusOK, goods, "")
	return
}

func Delete(c *gin.Context) {
	var CreateGoodsReq struct {
		ID int64
	}
	err := c.ShouldBind(&CreateGoodsReq)
	if err != nil {
		Res(c, 10000, nil, err.Error())
		return
	}
	err = service.DeleteGoods(c, CreateGoodsReq.ID)
	if err != nil {
		Res(c, 10001, nil, err.Error())
		return
	}
	Res(c, http.StatusOK, nil, "")
	return
}

func Update(c *gin.Context) {
	var UploadGoodsReq struct {
		ID         int64
		GoodsName  string
		GoodsPrice float64
		GoodsNum   int64
	}
	err := c.ShouldBind(&UploadGoodsReq)
	if err != nil {
		Res(c, 10000, nil, err.Error())
		return
	}
	goods, err := service.UpdateGoods(c, &proto.GoodsInfo{
		ID:         UploadGoodsReq.ID,
		GoodsName:  UploadGoodsReq.GoodsName,
		GoodsPrice: float32(UploadGoodsReq.GoodsPrice),
		GoodsNum:   UploadGoodsReq.GoodsNum,
	})
	if err != nil {
		return
	}
	if err != nil {
		Res(c, 10001, goods, err.Error())
		return
	}
	Res(c, http.StatusOK, goods, "")
	return
}

func Where(c *gin.Context) {
	var CreateGoodsReq struct {
		Name string
	}
	err := c.ShouldBind(&CreateGoodsReq)
	if err != nil {
		Res(c, 10000, nil, err.Error())
		return
	}
	goods, err := service.WhereGoods(c, CreateGoodsReq.Name)
	if err != nil {
		return
	}
	if err != nil {
		Res(c, 10001, goods, err.Error())
		return
	}
	Res(c, http.StatusOK, goods, "")
	return
}

func UploadFile(c *gin.Context) {
	var CreateGoodsReq struct {
		ID int64
	}
	err := c.ShouldBind(&CreateGoodsReq)
	if err != nil {
		Res(c, 10000, nil, err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		Res(c, 10001, nil, err.Error())
		return
	}
	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
	Res(c, http.StatusOK, nil, "")
	return
}

type Response struct {
	Code int64
	Data interface{}
	Msg  string
}

func Res(c *gin.Context, code int64, data interface{}, msg string) {
	httpCode := http.StatusOK
	if code > 20000 {
		httpCode = http.StatusBadGateway
	}
	c.JSON(httpCode, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	return
}
