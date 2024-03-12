package routers

import (
	c "2108-a-homework/my-app/controllers"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.New()
	Router.POST("login", c.Login)
}
