// Code generated by goctl. DO NOT EDIT.
package types

type UserLogin struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Tel      string `form:"tel"`
	Code     int64  `form:"code"`
}

type UserRegister struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Tel      string `form:"tel"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}