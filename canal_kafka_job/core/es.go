package core

import (
	"canal_kafka_job/global"
	"context"
	"github.com/olivere/elastic/v7"
)

func InitEs() {
	// 定义 Elasticsearch 服务器的地址
	elasticURL := "http://47.109.31.106:9200"
	// 创建 Elasticsearch 客户端
	client, err := elastic.NewClient(
		elastic.SetURL(elasticURL),
		elastic.SetSniff(false),      //elastic.SetSniff(false) 禁用节点发现功能,否则会报错
		elastic.SetBasicAuth("", ""), // 用户名密码
	)
	if err != nil {
		panic(err)
	}
	// 使用上下文进行 Ping 操作
	ctx := context.Background()
	_, _, err = client.Ping(elasticURL).Do(ctx)
	if err != nil {
		panic(err)
	}
	// 将对象放入全局变量
	global.ESClient = client
}
