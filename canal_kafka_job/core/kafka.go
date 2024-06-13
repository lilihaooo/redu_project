package core

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func InitKafkaConn(topic string, partition int) *kafka.Conn {
	// 连接至Kafka集群的Leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "47.109.31.106:9092", topic, partition)
	if err != nil {
		panic(err)
	}
	return conn
}
