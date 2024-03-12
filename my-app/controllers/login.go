package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "wanan" {
		c.JSON(0, gin.H{
			"msg": "登录失败,账号不存在",
		})
	}
	if password == "1104127326..." {
		c.JSON(0, gin.H{
			"code": http.StatusOK,
			"msg":  "登录失败,密码不正确",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登陆成功",
	})

}
