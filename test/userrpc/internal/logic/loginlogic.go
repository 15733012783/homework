package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"net/http"
	"test/server"
	"test/userrpc/model"

	"test/userrpc/internal/svc"
	"test/userrpc/userrpc"

	"github.com/zeromicro/go-zero/core/logx"

	_ "google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type User struct {
	Id       int
	Username string `valid:"Required;Match(/^Bee.*/)"`
	Password string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头 	// Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	Mobile   string `valid:"Mobile"`                   // Mobile 必须为正确的手机号 	// IP 必须为一个正确的 IPv4 地址
}

func Md5(password string) string {
	Password := []byte(password)
	sum := md5.Sum(Password)
	sprintf := fmt.Sprintf("%x", sum)
	return sprintf
}

func TableVerify(u User) bool {
	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		return false
	}
	return b
}

func returned(code int, msg string, data []byte) *userrpc.StreamResp {
	StreamResp := &userrpc.StreamResp{
		Code: int64(code),
		Msg:  msg,
		Data: data,
	}
	return StreamResp
}

func (l *LoginLogic) Login(in *userrpc.Login) (*userrpc.StreamResp, error) {
	table := TableVerify(User{Username: in.Username, Password: in.Password, Mobile: in.Tel})
	if !table {
		return returned(
			http.StatusInternalServerError, "登陆失败，数据格式不正确", nil,
		), nil
	}

	err, login := model.UserLogin(model.User{
		Username: in.Username,
		Password: in.Password,
		Tel:      in.Tel,
	})

	if err != nil {
		return returned(
			http.StatusInternalServerError, "登陆失败，账号不存在", nil,
		), nil
	}

	md5Password := Md5(in.Password)
	if login.Password != md5Password {
		return returned(
			http.StatusInternalServerError, "登陆失败，账号密码不正确", nil,
		), nil
	}

	token := server.SetToken(login.Username)
	return returned(
		http.StatusInternalServerError, "登陆成功", []byte(token),
	), nil
}
