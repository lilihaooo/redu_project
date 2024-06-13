package seckill

import (
	"github.com/jinzhu/copier"
	"redu/global"
	"redu/model"
	"redu/serialize/common"
	"redu/util"
)

type Seckill struct {
	ID             uint                 `json:"id"`
	ProductId      int                  `json:"product_id"`
	ProductName    string               `json:"product_name"`
	ShopName       string               `json:"shop_name"`
	Title          string               `json:"title"`
	BeginTime      string               `json:"begin_time"`
	EndTime        string               `json:"end_time"`
	Count          int                  `json:"count"`
	UnitPrice      float64              `json:"unit_price"`
	AllowUserGrade int                  `json:"allow_user_grade"`
	Status         common.SeckillStatus `json:"status"`
}

func BuildSeckill(m model.Seckill) (s Seckill) {
	// 使用 Copier 复制字段
	if err := copier.Copy(&s, m); err != nil {
		// 处理错误
		global.Logrus.Error(err)
		return
	}
	s.ProductName = m.Product.Title
	s.ShopName = m.Product.Shop.Name
	s.BeginTime = util.TimeToString(m.BeginTime, "2006-01-02 15:04:05")
	s.EndTime = util.TimeToString(m.EndTime, "2006-01-02 15:04:05")
	return s
}

func BuildSeckillList(models []model.Seckill) (list []Seckill) {
	for _, m := range models {
		one := BuildSeckill(m)
		list = append(list, one)
	}
	return
}
