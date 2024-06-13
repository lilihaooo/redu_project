package seckill

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"redu/global"
	"redu/model"
	"redu/request"
	"redu/serialize/common"
	"redu/util"
	"strconv"
	"time"
)

func (s Seckill) CreateSeckillActivity(param *request.CreateSeckillActivityParam) error {
	// 检查商品是否存在
	var productM model.Product
	productM.ID = param.ProductID
	product, err := productM.GetOneByID()
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	// 判断库存是否足够
	if product.Remaining < param.Count {
		return errors.New("库存不足")
	}

	// 判断金额是否设置合理
	if product.BottomPrice <= param.UnitPrice {
		return errors.New("单价过高")
	}

	var seckillM model.Seckill
	// 使用 Copier 复制字段
	if err = copier.Copy(&seckillM, param); err != nil {
		// 处理错误
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	seckillM.BeginTime = util.TimeParse(param.BeginTime, "2006-01-02 15:04:05")
	seckillM.EndTime = util.TimeParse(param.EndTime, "2006-01-02 15:04:05")

	fmt.Println(param.BeginTime)
	fmt.Println(seckillM.BeginTime)

	if err = seckillM.CreateOne(); err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	// 创建一个秒杀 redis计数器
	counterKey := "seckill:counter:" + strconv.Itoa(int(seckillM.ID))
	err = util.RedisStrSet(counterKey, strconv.Itoa(seckillM.Count), 0)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("redis计数器生成失败, 请联系管理员")
	}

	// 创建秒杀活动 将mysql中的数据保存到redis中
	seckillActivityKey := "seckill:activity:" + strconv.Itoa(int(seckillM.ID))
	seckillM.Product = nil // 不序列化到json对象中
	seckillM.Status = int(common.SeckillStatusNormal)
	valBite, err := json.Marshal(seckillM)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}
	err = util.RedisStrSet(seckillActivityKey, string(valBite), 0)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("创建活动失败, 请联系管理员")
	}

	err = global.DB.Table("seckill").Where("id = ?", seckillM.ID).Update("status", 1).Error
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("活动启动失败, 请联系管理员")
	}

	// 新增今日秒杀活动后,  删除今天的 redis t_list key
	now := time.Now()
	nowStr := util.TimeToString(now, "2006-01-02")
	tListKey := "seckill:today_list:" + nowStr
	_, err = global.RedisClient.Del(context.Background(), tListKey).Result()
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	return nil
}
