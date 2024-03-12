package api

import (
	"2108-a-homework/test5/api/goods"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	goods.Register(r)
}
