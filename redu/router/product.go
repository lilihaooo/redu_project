package router

import (
	"github.com/gin-gonic/gin"
	"redu/api/v1"
	"redu/api/v2"
)

func ProductRouter(appGroup *gin.RouterGroup) {
	// v1 商品列表
	v1ProductGroup := appGroup.Group("v1/product")
	v1ProductGroup.POST("list", v1.GroupApp.SearchProductList) // 替换为v2 es搜索版本

	// v2 商品列表
	v2ProductGroup := appGroup.Group("v2/product")
	v2ProductGroup.POST("list", v2.GroupApp.SearchProductList) // es搜索版本

	// 商品分类
	v1ProductGroup.GET("category/tree", v1.GroupApp.GetCategoryTree)

}
