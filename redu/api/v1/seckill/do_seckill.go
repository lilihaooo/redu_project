package seckill

import (
	"github.com/gin-gonic/gin"
	"redu/request"
	"redu/response"
	"redu/service"
)

func (Seckill) DoSeckill(c *gin.Context) {
	id := request.GetIDParam(c, "id")

	err := service.AppService.DoSeckill(id)
	if err != nil {
		response.FailWithMsg(c, response.OTHER_ERROR, err.Error())
		return
	}

	response.OkWithMsg(c, "")
}
