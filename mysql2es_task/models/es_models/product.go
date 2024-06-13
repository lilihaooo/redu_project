package es_models

import (
	"context"
	"github.com/olivere/elastic/v7"

	"mysql2es_task/global"
)

type Product struct {
	ID                uint   `json:"id"`
	XingxuanCredit    string `json:"xingxuan_credit"`
	ProductCategoryId string `json:"product_category_id"`
	SalesAll          string `json:"sales_all"`
	//Sales24           string  `json:"sales_24"` //  0   2.5w+
	Platform        string  `json:"platform"`
	Title           string  `json:"title"`
	Commission      string  `json:"commission"`
	BottomPrice     float64 `json:"bottom_price"`
	CommissionValue float64 `json:"commission_value"`
	Remaining       int     `json:"remaining"`
	ExperienceScore float64 `json:"experience_score"` // 体验分
	IsFake          int     `json:"is_fake"`          // 假数据
}

func (p Product) Index() string {
	return "product_index"
}

func (p Product) Mapping() string {
	return `
{
  "settings": {:
    "index":{
      "max_result_window": "1000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "id": { 
        "type": "integer"
      },
      "xingxuan_credit": { 
        "type": "keyword"
      },
      "product_category_id": {
        "type": "integer"
      },
      "sales_all": {
        "type": "integer"
      },
      "sales_24": {
        "type": "integer"
      },
      "commission": {
        "type": "integer"
      },
      "bottom_price": {
        "type": "float"
      },
      "platform": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "commission_value": { 
        "type": "double"
      },
      "remaining": { 
        "type": "integer"
      },
      "experience_score": {
        "type": "float"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (p Product) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(p.Index()).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (p Product) CreateIndex() error {
	if p.IndexExists() {
		// 有索引
		err := p.RemoveIndex()
		if err != nil {
			return err
		}
	}
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(p.Index()).
		BodyString(p.Mapping()).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return err
	}
	if !createIndex.Acknowledged {
		global.Logrus.Error("创建失败")
		return err
	}
	global.Logrus.Infof("索引 %s 创建成功", p.Index())
	return nil
}

// RemoveIndex 删除索引
func (p Product) RemoveIndex() error {
	global.Logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(p.Index()).Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return err
	}
	if !indexDelete.Acknowledged {
		global.Logrus.Error("删除索引失败")
		return err
	}
	global.Logrus.Info("索引删除成功")
	return nil
}

// ISExistData 是否存在该文章
func (p Product) ISExistData() bool {
	res, err := global.ESClient.
		Search(p.Index()).
		Query(elastic.NewTermQuery("keyword", p.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

//func (p Product) EsGetArticleList(pageReq *req.PaginationReq, key string) (list []Product, count int64, err error) {
//offset := req.GetOffset(pageReq)
//if offset < 0 {
//	offset = 0
//}
//boolQuery := elastic.NewBoolQuery()
//if key != "" {
//	boolQuery.Must(
//		elastic.NewMatchQuery("title", key),
//	)
//}
//searchResult, err := global.ESClient.Search().
//	Index(p.Index()).
//	Query(boolQuery).                    // 查询条件
//	From(offset).Size(pageReq.PageSize). // 分页
//	Do(context.Background())
//if err != nil {
//	return nil, 0, err
//}
//count = searchResult.Hits.TotalHits.Value
//if count == 0 {
//	return []Product{}, count, nil
//}
//for _, hit := range searchResult.Hits.Hits {
//	err = json.Unmarshal(hit.Source, &p)
//	if err != nil {
//		return nil, 0, err
//	}
//	p.ID = hit.Id
//	list = append(list, p)
//}
//return list, count, nil

//}
