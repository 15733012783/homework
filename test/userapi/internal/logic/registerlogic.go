package logic

import (
	"context"
	"test/userrpc/userrpc"

	"test/userapi/internal/svc"
	"test/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegister) (resp *types.Response, err error) {
	add, err := l.svcCtx.Userrpc.Register(l.ctx, &userrpc.UserRegister{
		Username: req.Username,
		Password: req.Password,
		Tel:      req.Tel,
	})

	return &types.Response{
		Code: int(add.Code),
		Msg:  add.Msg,
		Data: add.Data,
	}, err
}
