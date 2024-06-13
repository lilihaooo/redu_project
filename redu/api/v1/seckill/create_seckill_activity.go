package seckill

import (
	"github.com/gin-gonic/gin"
	"redu/global"
	"redu/request"
	"redu/response"
	"redu/service"
	"redu/util"
)

func (Seckill) CreateSeckillActivity(c *gin.Context) {
	var param request.CreateSeckillActivityParam
	if err := c.ShouldBindJSON(&param); err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	// 校验参数
	if err := util.ZhValidate(&param); err != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, err)
		return
	}
	// 校验时间是否合规
	if param.EndTime <= param.BeginTime {
		response.FailWithMsg(c, response.FAIL_VALIDATE, "日期设置有误")
		return
	}
	if err := service.AppService.CreateSeckillActivity(&param); err != nil {
		response.FailWithMsg(c, response.OTHER_ERROR, err.Error())
		return
	}

	response.OkWithData(c, param)
}
