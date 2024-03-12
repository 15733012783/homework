package goods

import (
	middleware "2108-a-homework/test5/jwt"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/create", Create)
	r.POST("/delete", Delete)
	r.POST("/update", Update)
	r.POST("/where", Where)
	u := r.Group("/file")
	u.Use(middleware.AuthMiddleware)
	{
		u.POST("info", UploadFile)
	}
}
