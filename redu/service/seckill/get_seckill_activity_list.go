package seckill

import (
	"errors"
	"redu/global"
	"redu/model"
	"redu/request"
	"redu/util"
)

func (s Seckill) GetSeckillActivityList(param *request.GetSeckillListParam) (list []model.Seckill, err error) {
	offset := util.GetOffset(param.Page, param.PageSize)
	seckills := []model.Seckill{}
	err = global.DB.Preload("Product").Preload("Product.Shop").Order("created_at desc").Limit(param.PageSize).Offset(offset).Find(&seckills).Error
	if err != nil {
		global.Logrus.Error(err)
		return seckills, errors.New("内部错误")
	}
	return seckills, nil
}
