package seckill

import (
	"github.com/gin-gonic/gin"
	"redu/global"
	"redu/request"
	"redu/response"
	"redu/serialize/seckill"
	"redu/service"
	"redu/util"
)

func (Seckill) GetTodaySeckillActivityList(c *gin.Context) {
	var param request.GetSeckillListParam

	if err := c.ShouldBindQuery(&param); err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.INVALID_PARAMS, "")
		return
	}
	// 校验参数
	if err := util.ZhValidate(&param); err != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, err)
		return
	}
	if param.PageSize == 0 {
		param.PageSize = 10
	}
	if param.Page == 0 {
		param.Page = 1
	}
	models, err := service.AppService.GetTodaySeckillActivityList(&param)
	if err != nil {
		response.FailWithMsg(c, response.OTHER_ERROR, err.Error())
		return
	}
	list := seckill.BuildSeckillList(models)
	response.OkWithData(c, list)
}
