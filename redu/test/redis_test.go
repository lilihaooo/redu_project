package test

import (
	"fmt"
	"redu/core"
	"redu/global"
	"redu/util"
	"testing"
)

func TestRedis(t *testing.T) {

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

	// string 计数器
	err := util.RedisStrSet("count:test", "5", 0)
	if err != nil {
		global.Logrus.Error(err)
		return
	}
	//in := global.RedisClient.Incr(context.Background(), "count_test")
	//fmt.Println(in)
	//global.Logrus.Error("eee")
	fmt.Println("test ok")

}
