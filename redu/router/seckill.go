package router

import (
	"github.com/gin-gonic/gin"

	"redu/api/v1"
)

func SeckillRouter(appGroup *gin.RouterGroup) {
	v1SeckillGroup := appGroup.Group("v1/seckill_activity")
	v1SeckillGroup.POST("", v1.GroupApp.CreateSeckillActivity)
	v1SeckillGroup.PUT("", v1.GroupApp.UpdateSeckillActivity)

	v1SeckillGroup.GET("list", v1.GroupApp.GetSeckillActivityList)

	v1SeckillGroup.GET("t_list", v1.GroupApp.GetTodaySeckillActivityList)
	v1SeckillGroup.DELETE("", v1.GroupApp.DeleteSeckillActivity)

	v1SeckillGroup.GET("do", v1.GroupApp.DoSeckill)

}
