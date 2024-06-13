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

func (s Seckill) UpdateSeckillActivity(param *request.UpdateSeckillActivityParam) error {
	// 判断活动是否存在
	var seckillM model.Seckill
	seckillM.ID = param.ID
	seckill, err := seckillM.GetOneByID()
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	// 使用 Copier 复制字段
	if err = copier.Copy(&seckill, param); err != nil {
		// 处理错误
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	seckill.BeginTime = util.TimeParse(param.BeginTime, "2006-01-02 15:04:05")
	seckill.EndTime = util.TimeParse(param.EndTime, "2006-01-02 15:04:05")

	fmt.Println(seckill.BeginTime)
	fmt.Println(seckill.EndTime)

	err = global.DB.Updates(&seckill).Error
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	// 创建一个秒杀 redis计数器
	counterKey := "seckill:counter:" + strconv.Itoa(int(seckill.ID))
	err = util.RedisStrSet(counterKey, strconv.Itoa(seckill.Count), 0)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("redis计数器生成失败, 请联系管理员")
	}

	// 覆盖秒杀活动 将mysql中的数据保存到redis中
	seckillActivityKey := "seckill:activity:" + strconv.Itoa(int(seckill.ID))
	seckill.Product = nil // 不序列化到json对象中
	seckill.Status = int(common.SeckillStatusNormal)
	valBite, err := json.Marshal(seckill)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}
	err = util.RedisStrSet(seckillActivityKey, string(valBite), 0)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("修改活动失败, 请联系管理员")
	}

	// 新增秒杀活动后,  删除今天的 redis t_list key
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
