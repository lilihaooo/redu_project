package product

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"redu/global"
	"redu/model"
	"redu/request"
	"redu/util"
	"strconv"
)

func (p Product) GetSearchProductsByEs(req *request.SearchProductListParam) (list []model.Product, count int64, err error) {
	// 构建 Elasticsearch 查询
	boolQuery := elastic.NewBoolQuery()

	// 不是fake数据
	if req.NotFake {
		boolQuery = boolQuery.Must(elastic.NewTermQuery("is_fake", 0))
	}

	// 平台分类
	boolQuery = boolQuery.Must(elastic.NewTermQuery("platform", req.Platform))

	// tag
	if len(req.ProductTagIds) > 0 {
		terms := make([]interface{}, len(req.ProductTagIds))
		for i, id := range req.ProductTagIds {
			terms[i] = id
		}
		boolQuery = boolQuery.Filter(elastic.NewTermsQuery("tag_id", terms...))
	}

	// 商品分类
	if req.ProductCategoryId != "" {
		// 获得分类信息
		var categoryM model.Category
		err = global.DB.Take(&categoryM, req.ProductCategoryId).Error
		if err != nil {
			global.Logrus.Error(err)
			return nil, 0, err
		}

		var categoryIDs []string

		if categoryM.Level == "1" {
			categoryIDs = p.GetIDsByLevel1(req.ProductCategoryId)
			// 将 []string 转换为 []interface{}
			var interfaceIds []interface{}
			for _, id := range categoryIDs {
				interfaceIds = append(interfaceIds, id)
			}
			boolQuery = boolQuery.Filter(elastic.NewTermsQuery("product_category_id", interfaceIds...))
		} else if categoryM.Level == "2" {
			categoryIDs = p.GetIDsByLevel2(req.ProductCategoryId)
			var interfaceIds []interface{}
			for _, id := range categoryIDs {
				interfaceIds = append(interfaceIds, id)
			}
			boolQuery = boolQuery.Filter(elastic.NewTermsQuery("product_category_id", interfaceIds...))
		} else if categoryM.Level == "3" {
			boolQuery = boolQuery.Filter(elastic.NewTermQuery("product_category_id", req.ProductCategoryId))
		}
	}

	// 关键词搜索
	//if req.Kw != "" {
	//	boolQuery = boolQuery.Must(elastic.NewMatchQuery("title", req.Kw))
	//}

	// 关键词搜索 - 模糊匹配
	if req.Kw != "" {
		boolQuery = boolQuery.Must(elastic.NewMatchPhraseQuery("title", req.Kw).Slop(0))
	}

	// 佣金比例筛选
	if req.Cm != "" {
		boolQuery = boolQuery.Filter(elastic.NewRangeQuery("commission").Gte(req.Cm))
	}

	// 到手价筛选
	if req.Pp0 != "" {
		boolQuery = boolQuery.Filter(elastic.NewRangeQuery("bottom_price").Gte(req.Pp0))
	}
	if req.Pp1 != "" {
		boolQuery = boolQuery.Filter(elastic.NewRangeQuery("bottom_price").Lte(req.Pp1))
	}

	// 销量筛选
	if req.F0 != "" {
		boolQuery = boolQuery.Filter(elastic.NewRangeQuery("sales_all").Gte(req.F0))
	}
	if req.F1 != "" {
		boolQuery = boolQuery.Filter(elastic.NewRangeQuery("sales_all").Lte(req.F1))
	}

	// 查询出总数量
	count, err = global.ESClient.Count("product_index").Query(boolQuery).Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return nil, 0, err
	}
	// 设置了最大返回数量
	if count > 1000 {
		count = 1000
	}

	// 创建搜索请求
	offset := util.GetOffset(req.Page, request.PageSize)
	searchRequest := global.ESClient.Search().Index("product_index").Query(boolQuery).From(offset).Size(request.PageSize)

	// 按照佣金比例排序
	if req.By == "commission" {
		searchRequest = searchRequest.Sort("commission", false)
	} else if req.By == "commission_value" {
		// 按照佣金价格排序
		searchRequest = searchRequest.Sort("commission_value", false)
	} else if req.By == "sales_all" {
		// 按照总销量排序
		searchRequest = searchRequest.Sort("sales_all", false)
	} else if req.By == "twenty_four_hour_sales" {
		// 按照近24小时销量排序
		searchRequest = searchRequest.Sort("sales_24", false)
	} else if req.By == "bottom_price" && req.Sort == "ASC" {
		// 按照到手价升序
		searchRequest = searchRequest.Sort("bottom_price", true)
	} else if req.By == "bottom_price" && req.Sort == "DESC" {
		// 按照到手价降序
		searchRequest = searchRequest.Sort("bottom_price", false)
	} else {
		searchRequest = searchRequest.Sort("id", false)
	}

	searchRequest = searchRequest.FetchSourceContext(elastic.NewFetchSourceContext(true).Include("id"))

	// 执行搜索请求
	searchResult, err := searchRequest.Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return nil, 0, err
	}

	type R struct {
		ID interface{} `json:"id"`
	}
	// 解析结果
	var IDs []int64
	for _, hit := range searchResult.Hits.Hits {
		var r R
		_ = json.Unmarshal(hit.Source, &r)
		var idInt64 int64
		switch v := r.ID.(type) {
		case float64: // JSON numbers are parsed as float64
			idInt64 = int64(v)
		case string:
			idInt64, err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				global.Logrus.Error(err)
			}
		}
		IDs = append(IDs, idInt64)
	}

	// 重新排序  (捞出来后重新排序)
	db := global.DB
	// 按照佣金比例排序
	if req.By == "commission" {
		db = db.Order("commission desc")
	}
	// 按照佣金价格排序
	if req.By == "commission_value" {
		db = db.Order("commission_value desc")
	}

	// 按照总销量排序
	if req.By == "sales" {
		db = db.Order("sales_all desc")
	}
	// 按照30日销量排序
	//if req.By == "thirty_day_sales" {
	//	db = db.Order("thirty_day_sales desc")
	//}
	// 按照近24小时量排序
	if req.By == "twenty_four_hour_sales" {
		db = db.Order("sales24 desc")
	}
	// 按照总销量排序
	if req.By == "sales_all" {
		db = db.Order("sales_all desc")
	}

	// 按照手价排序
	if req.By == "bottom_price" {
		if req.Sort == "DESC" {
			db = db.Order("bottom_price desc")
		}
		if req.Sort == "ASC" {
			db = db.Order("bottom_price asc")
		}
	}

	err = db.Where("id in ?", IDs).
		Preload("Shop").
		Preload("ProductImages").
		Find(&list).Error
	if err != nil {
		global.Logrus.Error(err)
		return nil, 0, err
	}
	return
}
