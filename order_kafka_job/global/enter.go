package global

import (
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var (
	Logger    *logrus.Logger
	KafkaConn *kafka.Conn
)
