package seckill

import (
	"encoding/json"
	"errors"
	"redu/global"
	"redu/model"
	"redu/request"
	"redu/util"
	"time"
)

func (s Seckill) GetTodaySeckillActivityList(param *request.GetSeckillListParam) (list []model.Seckill, err error) {
	seckills := []model.Seckill{}
	date := util.TimeToString(time.Now(), "2006-01-02")
	// 从redis中获取列表
	key := "seckill:today_list:" + date
	// 通过 key 获取 value
	val, err := util.RedisStrGet(key)
	if err == nil {
		// 查询到了, 将json解析到结构体
		err = json.Unmarshal([]byte(val), &seckills)
		if err != nil {
			global.Logrus.Error(err)
		} else {
			return seckills, nil
		}
	}

	offset := util.GetOffset(param.Page, param.PageSize)
	err = global.DB.
		Preload("Product").
		Preload("Product.Shop").
		Where("DATE(begin_time) = ?", date).
		Order("begin_time asc").
		Limit(param.PageSize).
		Offset(offset).
		Find(&seckills).
		Error
	if err != nil {
		global.Logrus.Error(err)
		return seckills, errors.New("内部错误")
	}

	// 加入redis中防止并发刷新网页获取列表
	valBite, err := json.Marshal(seckills)
	if err != nil {
		global.Logrus.Error(err)
	}
	expireDuration := 10 * time.Minute
	err = util.RedisStrSet(key, string(valBite), expireDuration)
	if err != nil {
		global.Logrus.Error(err)
	}
	return seckills, nil
}
