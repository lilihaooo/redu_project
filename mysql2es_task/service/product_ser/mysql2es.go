package product_ser

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"mysql2es_task/global"
	"mysql2es_task/models/es_models"
	"strconv"
	"time"
)

func Mysql2Es() {

	//创建索引
	//var model es_models.Product
	//err := model.CreateIndex()
	//if err != nil {
	//	fmt.Println("失败")
	//} else {
	//	fmt.Println("更新索引成功")
	//	return
	//}

	a := time.Now()

	var offset = 0
	for {
		// 从数据库中获取100条记录
		var products []es_models.Product
		sql := "SELECT b.id, b.xingxuan_credit, b.product_category_id, b.sales_all, b.platform, b.title, b.commission, b.bottom_price, b.commission_value, b.remaining, b.experience_score, b.is_fake FROM product b RIGHT JOIN (SELECT id FROM product ORDER BY id ASC LIMIT " + strconv.Itoa(offset) + ", 10000) a ON a.id = b.id"
		err := global.DB.Raw(sql).Scan(&products).Error
		if err != nil {
			global.Logrus.Error(err)
			return
		}
		count := len(products)
		if count == 0 {
			b := time.Since(a)
			fmt.Println(b)
			break
		}
		toEs(products)
		offset += count

		global.ESClient.Flush()
		// 使用Count API来获取索引中的总记录数量
		total, _ := global.ESClient.Count("product_index").Do(context.Background())
		// 打印总记录数量
		fmt.Printf("总记录数量为：%d\n", total)
	}
	// count = 10000256
}

func toEs(products []es_models.Product) {
	var model es_models.Product

	// note 可以将商品id指定为文档id
	// 构建批量请求
	bulkRequest := global.ESClient.Bulk()
	for _, item := range products {

		indexReq := elastic.NewBulkIndexRequest().Index(model.Index()).Doc(item)
		bulkRequest = bulkRequest.Add(indexReq)
	}

	// 执行批量请求
	bulkResponse, err := bulkRequest.Refresh("true").Do(context.Background()) // Refresh("true")立刻刷新
	if err != nil {
		global.Logrus.Error(err)
		return
	}

	fmt.Println(len(bulkResponse.Succeeded()))

	// 处理批量响应
	if bulkResponse.Errors {
		// 处理错误
		for _, failedItem := range bulkResponse.Failed() {
			fmt.Printf("Failed to index document %s: %s", failedItem.Id, failedItem.Error)
		}
	}
}
