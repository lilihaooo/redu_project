package router

import (
	"github.com/gin-gonic/gin"
	"redu/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode("release")
	r := gin.Default()
	//r.POST("/api/login", v1.ApiGroupApp.UserApi.Login)
	// 需要登陆
	//r.Use(jwt.JwtAuth())
	r.Use(middleware.RequestCostHandler())
	apiGroup := r.Group("api")

	ProductRouter(apiGroup)
	SeckillRouter(apiGroup)

	return r
}
