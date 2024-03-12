package logic

import (
	"context"
	"test/userrpc/userrpc"

	"test/userapi/internal/svc"
	"test/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLogin) (resp *types.Response, err error) {
	add, err := l.svcCtx.Userrpc.Login(l.ctx, &userrpc.Login{
		Username: req.Username,
		Password: req.Password,
		Tel:      req.Tel,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Code: int(add.Code),
		Msg:  add.Msg,
		Data: string(add.Data),
	}, err
}
