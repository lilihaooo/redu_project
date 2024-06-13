package main

import (
	"redu/core"
	"redu/flag"
	"redu/global"
	"redu/router"
)

func main() {
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

	// 控制台操作
	option := flag.Parse()
	flag.SwitchOption(option)
	if flag.IsWebStop(option) {
		return
	}

	r := router.InitRouter()
	addr := global.Config.Server.Addr()
	global.Logrus.Infof("http server 运行在: %s", addr)
	err := r.Run(addr)
	if err != nil {
		global.Logrus.Fatal(err.Error())
	}
}
