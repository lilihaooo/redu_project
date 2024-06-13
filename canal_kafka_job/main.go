package main

import (
	"canal_kafka_job/core"
	"canal_kafka_job/job"
)

func main() {
	core.InitLogrus()
	core.InitEs()

	job.Kafka2ES()
}
