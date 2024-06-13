package job

import (
	"canal_kafka_job/global"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func Kafka2ES() {
	// 创建一个reader，指定GroupID，从 topic-A 消费消息
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"47.109.31.106:9092"},
		GroupID:  "consumer-group-id", // 指定消费者组id
		Topic:    "topic_redu",
		MaxBytes: 10e6, // 10MB
	})

	fmt.Println("连接成功...")

	// 接收消息
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		var msg ProductMassage
		err = json.Unmarshal(m.Value, &msg)
		if err != nil {
			global.Logger.Error(err)
			continue
		}
		esProduct := msg.Data[0]
		// 判断消息的类型
		if msg.Type == "UPDATE" {
			// 获取文档_id
			_id := esProduct.GetDocIDByProductID()
			err = esProduct.EsUpdateProduct(_id)
			if err != nil {
				global.Logger.Error(err)
				continue
			}
			fmt.Println("kafka to es 修改成功")
			continue

		}
		if msg.Database != "redu" || msg.Table != "product" {
			continue
		}
		if msg.Type == "INSERT" {
			err = esProduct.EsCreateProduct()
			if err != nil {
				global.Logger.Error(err)
				continue
			}
			fmt.Println("kafka to es 添加成功")
			continue
		}

		if msg.Type == "DELETE" {
			// 获取文档_id
			_id := esProduct.GetDocIDByProductID()
			err = esProduct.EsDeleteProduct(_id)
			if err != nil {
				global.Logger.Error(err)
				continue
			}
			fmt.Println("kafka to es 删除成功")
			continue
		}

	}
	// 程序退出前关闭Reader
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
