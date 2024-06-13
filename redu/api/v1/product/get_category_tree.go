package product

import (
	"github.com/gin-gonic/gin"
	"redu/response"
	"redu/service"
)

func (Product) GetCategoryTree(c *gin.Context) {
	tree, err := service.AppService.MakeCategoryTree()
	if err != nil {
		response.FailWithMsg(c, response.OTHER_ERROR, "")
		return
	}
	response.OkWithData(c, tree)
}
