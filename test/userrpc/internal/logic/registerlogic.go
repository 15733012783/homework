package logic

import (
	"context"
	"net/http"
	"test/userrpc/model"

	"test/userrpc/internal/svc"
	"test/userrpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *userrpc.UserRegister) (*userrpc.StreamResp, error) {
	table := TableVerify(User{Username: in.Username, Password: in.Password, Mobile: in.Tel})
	if !table {
		return &userrpc.StreamResp{
			Code: http.StatusInternalServerError,
			Msg:  "登陆失败，数据格式不正确",
			Data: nil,
		}, nil
	}

	login := model.UserCreate(in.Username)
	if login != nil {
		return &userrpc.StreamResp{
			Code: 0,
			Msg:  "账号已经被注册",
			Data: nil,
		}, nil
	}

	md5Password := Md5(in.Password)
	register := model.UserRegister(model.User{
		Username: in.Username,
		Password: md5Password,
		Tel:      in.Tel,
	})

	if register == 0 {
		return &userrpc.StreamResp{
			Code: http.StatusInternalServerError,
			Msg:  "注册失败",
			Data: nil,
		}, nil
	}

	return &userrpc.StreamResp{
		Code: http.StatusOK,
		Msg:  "注册成功",
		Data: nil,
	}, nil
}
