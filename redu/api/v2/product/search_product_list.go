package product

import (
	"github.com/gin-gonic/gin"
	"redu/global"
	"redu/request"
	"redu/response"
	"redu/service"
	"redu/util"
)

func (Product) SearchProductList(c *gin.Context) {
	var req request.SearchProductListParam
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.INVALID_PARAMS, err.Error())
		return
	}
	// 校验参数
	if err := util.ZhValidate(&req); err != "" {
		response.FailWithMsg(c, response.FAIL_VALIDATE, err)
		return
	}
	data, count, err := service.AppService.GetSearchProductsByEs(&req)
	if err != nil {
		response.FailWithMsg(c, response.ERROR_DB_OPER, "查询失败")
		return
	}
	response.OkWithList(c, data, count)
}
