package core

import (
	"context"
	"github.com/segmentio/kafka-go"
	"order_kafka_job/global"
)

func InitKafkaConn(topic string, partition int) {
	// 连接至Kafka集群的Leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "47.109.31.106:9092", topic, partition)
	if err != nil {
		panic(err)
	}
	global.KafkaConn = conn
}
