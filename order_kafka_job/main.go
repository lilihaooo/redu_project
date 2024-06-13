package main

import (
	"order_kafka_job/core"
	"order_kafka_job/job"
)

func main() {
	// 连接
	core.InitLogger()
	//core.InitKafkaConn("test", 0)
	job.Work()
}
