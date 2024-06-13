package seckill

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"redu/global"
	"redu/model"
	"redu/serialize/common"
	"redu/util"
	"strconv"
	"time"
)

func (s Seckill) DoSeckill(seckillID uint) error {
	// 获取秒杀活动
	var seckill model.Seckill
	seckillActivityKey := "seckill:activity:" + strconv.Itoa(int(seckillID))
	val, err := util.RedisStrGet(seckillActivityKey)
	if err != nil {
		// 查询出错
		global.Logrus.Error(err)
		return errors.New("内部错误")

	}
	// 通过es 查询到了结果
	err = json.Unmarshal([]byte(val), &seckill)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	// 判断时间是否在开始时间段
	now := time.Now()
	nowStr := util.TimeToString(now, "2006-01-02 15:04:05")
	bStr := util.TimeToString(seckill.BeginTime, "2006-01-02 15:04:05")
	eStr := util.TimeToString(seckill.EndTime, "2006-01-02 15:04:05")
	if bStr > nowStr {
		return errors.New("秒杀未开始")
	}
	if eStr < nowStr {
		return errors.New("秒杀已结束")
	}
	// 判断库存
	if seckill.Count == 0 {
		return errors.New("已经抢完")
	}
	// 判断状态
	if seckill.Status == int(common.SeckillStatusStop) {
		return errors.New("活动异常")

	}

	// 获取用户信息(userID 使用uuid模拟)
	userID := uuid.New()

	// 用户等级判断 todo

	counterKey := "seckill:counter:" + strconv.Itoa(int(seckill.ID))
	// 获取剩余数量
	numStr, err := util.RedisStrGet(counterKey)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}
	// 字符串数字转为int
	num, err := strconv.Atoi(numStr) // Atoi 是 ParseInt 的一个简化版，用于将字符串转换为 int 类型
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}
	if num <= 0 {
		return errors.New("已抢空")
	}

	// 减库存
	_, err = global.RedisClient.Decr(context.Background(), counterKey).Result()
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	/*模拟创建订单*/
	var order model.Order
	// 生成订单号
	orderNum := uuid.New()

	order.CreatedAt = time.Now()
	order.OrderNum = orderNum.String()
	order.UserID = 3
	order.OrderStatus = 1
	order.PayStatus = 1
	order.PayMethod = 1
	pNow := time.Now()
	order.PayAt = &pNow
	order.TotalAmount = 30
	order.CompletedAt = nil
	order.CancelledAt = nil
	order.Note = ""
	//order.OrderItems = 

	// LPUSH 操作，将元素推入列表的左端
	var values []string
	values = append(values, userID.String())
	orderKdy := "seckill:order:" + strconv.Itoa(int(seckill.ID))
	err = util.RedisLPush(orderKdy, values)
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	return nil
}
