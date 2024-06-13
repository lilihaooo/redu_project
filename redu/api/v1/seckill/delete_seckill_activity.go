package seckill

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"redu/global"
	"redu/model"
	"redu/request"
	"redu/response"
	"redu/util"
	"strconv"
	"time"
)

func (Seckill) DeleteSeckillActivity(c *gin.Context) {
	id := request.GetIDParam(c, "id")
	var seckillM model.Seckill
	seckillM.ID = id

	seckill, err := seckillM.GetOneByID()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailWithMsg(c, response.ERROR_NOT_EXIST_RECODE, "")
			return
		} else {
			global.Logrus.Error(err)
			response.FailWithMsg(c, response.OTHER_ERROR, "内部错误")
			return
		}
	}
	// 删除redis记录 活动计数器, 活动记录, 当天的列表id
	redisC := global.RedisClient
	// 删除计数器
	counterKey := "seckill:counter:" + strconv.Itoa(int(seckill.ID))
	_, err = redisC.Del(context.Background(), counterKey).Result()
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.OTHER_ERROR, "内部错误")
		return
	}
	// 删除 活动
	activityKey := "seckill:activity:" + strconv.Itoa(int(seckill.ID))
	_, err = redisC.Del(context.Background(), activityKey).Result()
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.OTHER_ERROR, "内部错误")
		return
	}
	// 删除 当天的列表
	date := util.TimeToString(time.Now(), "2006-01-02")
	TListKey := "seckill:today_list:" + date
	_, err = redisC.Del(context.Background(), TListKey).Result()
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.OTHER_ERROR, "内部错误")
		return
	}

	err = global.DB.Delete(&seckillM).Error
	if err != nil {
		global.Logrus.Error(err)
		response.FailWithMsg(c, response.OTHER_ERROR, "")
		return
	}

	response.OkWithMsg(c, "")
}
