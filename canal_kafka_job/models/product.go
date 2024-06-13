package models

import (
	"canal_kafka_job/global"
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

type Product struct {
	ID                uint   `json:"id,string"`
	XingxuanCredit    string `json:"xingxuan_credit"`
	ProductCategoryId string `json:"product_category_id"`
	SalesAll          string `json:"sales_all"`
	//Sales24           string  `json:"sales_24"` //  0   2.5w+
	Platform        string  `json:"platform"`
	Title           string  `json:"title"`
	Commission      string  `json:"commission"`
	BottomPrice     float64 `json:"bottom_price,string"`
	CommissionValue float64 `json:"commission_value,string"`
	Remaining       int     `json:"remaining,string,string"`
	ExperienceScore float64 `json:"experience_score,string"` // 体验分
	IsFake          int     `json:"is_fake,string"`          // 假数据
}

func (p *Product) Index() string {
	return "product_index"
}

func (p *Product) Mapping() string {
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
func (p *Product) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(p.Index()).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (p *Product) CreateIndex() error {
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
		global.Logger.Error(err)
		return err
	}
	if !createIndex.Acknowledged {
		global.Logger.Error("创建失败")
		return err
	}
	global.Logger.Infof("索引 %s 创建成功", p.Index())
	return nil
}

// RemoveIndex 删除索引
func (p *Product) RemoveIndex() error {
	global.Logger.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(p.Index()).Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	if !indexDelete.Acknowledged {
		global.Logger.Error("删除索引失败")
		return err
	}
	global.Logger.Info("索引删除成功")
	return nil
}

// ISExistData 是否存在该文章
func (p *Product) ISExistData() bool {
	res, err := global.ESClient.
		Search(p.Index()).
		Query(elastic.NewTermQuery("keyword", p.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (p *Product) EsCreateProduct() (err error) {
	_, err = global.ESClient.Index().
		Index(p.Index()).
		BodyJson(p).
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
	}
	return
}

func (p *Product) EsUpdateProduct(_id string) (err error) {
	// 全部字段更新
	_, err = global.ESClient.Index().
		Index(p.Index()).
		Id(_id). // 指定要更新的文档ID
		BodyJson(p).
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
	}
	return
}

func (p *Product) EsDeleteProduct(_id string) (err error) {
	_, err = global.ESClient.Delete().
		Index(p.Index()).
		Id(_id). // 指定要更新的文档ID
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
	}
	return
}

func (p *Product) GetDocIDByProductID() string {
	termQuery := elastic.NewTermQuery("id", p.ID)
	searchResult, err := global.ESClient.Search().
		Index(p.Index()).
		Query(termQuery).
		Do(context.Background())
	if err != nil {
		global.Logger.Error(err)
		return ""
	}
	// 处理搜索结果
	if searchResult.TotalHits() > 0 {
		return searchResult.Hits.Hits[0].Id
	}
	return ""
}
