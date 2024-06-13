package core

import (
	"canal_kafka_job/global"
	"github.com/sirupsen/logrus"
	"os"
)

type CustomHook struct {
	file *os.File
}

// InitLogrus 初始化日志
func InitLogrus() {
	// 创建新的Logger实例
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel) // debug级别 这意味着日志记录器会记录所有级别的日志消息，包括调试、信息、警告和错误等级别的消息。
	log.SetReportCaller(true)       // 是否显示行号
	formatter := &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:03:04",
		FullTimestamp:   true,
	}
	log.SetFormatter(formatter)
	// 设置日志输出的最低级别为DebugLevel
	global.Logger = log
}
