package request

import (
	"github.com/gin-gonic/gin"
	"redu/response"
	"strconv"
)

type PaginationParam struct {
	Page     int `form:"page" json:"page" label:"页码"`
	PageSize int `form:"page_size" json:"page_size" validate:"oneof=0 5 10 20 30" label:"每页条数"`
}

func GetIDParam(c *gin.Context, idStr string) uint {
	IDStr := c.Query(idStr)
	if IDStr == "" || IDStr == "0" {
		response.FailWithMsg(c, response.INVALID_PARAMS, idStr+"不能为空")
		return 0
	}
	ID, _ := strconv.Atoi(IDStr)
	return uint(ID)
}
