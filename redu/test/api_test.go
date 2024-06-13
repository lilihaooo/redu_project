package test

import (
	"fmt"
	"redu/core"
	"redu/request"
	"redu/service"
	"testing"
)

func TestApi(t *testing.T) {

	// 初始化配置文件
	core.InitConf()
	// 初始化日志
	core.InitLogrus()
	defer core.CloseLogFile()
	// 初始化Gorm
	core.InitGorm()
	defer core.CloseGormLogFile()
	// 初始化错误码json文件
	core.InitResMap()
	//初始化redis
	core.InitRedis()
	// 初始化es
	core.InitEs()

	param := request.GetSeckillListParam{}
	param.Page = 1
	param.PageSize = 10
	list, err := service.AppService.GetTodaySeckillActivityList(&param)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(list)
}
