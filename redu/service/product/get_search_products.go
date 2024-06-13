package product

import (
	"redu/global"
	"redu/model"
	"redu/request"
	"redu/util"
)

func (p Product) GetSearchProducts(req *request.SearchProductListParam) (list []model.Product, count int64, err error) {
	var productM model.Product
	db := global.DB.Model(productM)

	// 平台分类
	db = db.Where("platform = ?", req.Platform)

	// tag
	var productIDs []uint
	if len(req.ProductTagIds) > 0 {
		err = global.DB.Table("product_tag").
			Select("product_id").
			Where("tag_id IN ?", req.ProductTagIds).
			Group("product_id").
			Having("COUNT(DISTINCT tag_id) = ?", len(req.ProductTagIds)).
			Pluck("product_id", &productIDs).Error
		if err != nil {
			global.Logrus.Error(err)
			return nil, 0, err
		}
		db.Where("id in ?", productIDs)
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
		}
		if categoryM.Level == "2" {
			categoryIDs = p.GetIDsByLevel2(req.ProductCategoryId)
		}
		if categoryM.Level == "3" {
			db = db.Where("product_category_id = ?", req.ProductCategoryId)
		} else {
			//fmt.Println(categoryIDs)
			db = db.Where("product_category_id in ?", categoryIDs)
		}

	}

	// 关键词搜索
	if req.Kw != "" {
		db = db.Where("title like ?", "%"+req.Kw+"%")
	}

	// 按照上架时间排序
	//if req.By == "begin_time" {
	//	if req.Sort == "DESC" {
	//		db = db.Order("bottom_price desc")
	//	}
	//	if req.Sort == "ASC" {
	//		db = db.Order("bottom_price asc")
	//	}
	//}

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
		db = db.Order("sales desc")
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

	//筛选:
	// 佣金比例 参数: cm, 字段:commission
	if req.Cm != "" {
		db = db.Where("commission >= ?", req.Cm)
	}
	// 到手价 >= 参数:pp0, 字段: bottom_price
	if req.Pp0 != "" {
		db = db.Where("bottom_price >= ?", req.Pp0)
	}
	// 到手价 <= 参数:pp1, 字段: bottom_price
	if req.Pp1 != "" {
		db = db.Where("bottom_price <= ?", req.Pp1)
	}
	// 销量  >= 参数:f0, 字段:sales_all
	if req.F0 != "" {
		db = db.Where("sales_all >= ?", req.F0)
	}
	// 销量  <= 参数:f1, 字段:sales_all
	if req.F1 != "" {
		db = db.Where("sales_all <= ?", req.F1)
	}
	// 体验分 >= 参数:pf, 字段:shop_creadit_product
	//if req.F0 != "" {
	//	db = db.Where("sales_all >= ?", req.F0)
	//}

	//err = db.Count(&count).Error
	//if err != nil {
	//	global.Logrus.Error(err)
	//	return nil, 0, err
	//}

	//db.Preload("Shop").Preload("ProductImages").Where("id IN (?)", db.Table("product_tags").Select("product_id").Where("tag_id IN (?)", []int{1, 2}).SubQuery()).Find(&products)

	offset := util.GetOffset(req.Page, request.PageSize)

	// 覆盖索引法则
	var IDs []int64
	//sql := "SELECT Id FROM product  ORDER BY id DESC LIMIT " + strconv.Itoa(offset) + ", 100"
	//err = db.Raw("SELECT Id FROM product  ORDER BY id DESC LIMIT ?, ?", offset, product_req.PageSize).Scan(&IDs).Error

	err = db.Select("id").Limit(request.PageSize).Offset(offset).Order("id DESC").Scan(&IDs).Error
	if err != nil {
		global.Logrus.Error(err)
		return nil, 0, err
	}

	count = 2000
	//err = db.Order("id desc").
	//	Limit(product_req.PageSize).
	//	Offset(offset).
	//	Preload("Shop").
	//	Preload("ProductImages").
	//	Find(&list).Error

	err = global.DB.Order("id desc").Where("id in ?", IDs).
		Preload("Shop").
		Preload("ProductImages").
		Find(&list).Error
	if err != nil {
		global.Logrus.Error(err)
		return nil, 0, err
	}
	return
}
