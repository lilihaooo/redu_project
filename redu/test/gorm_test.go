package test

import (
	"fmt"
	"redu/core"
	"redu/global"
	"redu/model"
	"testing"
	"time"
)

func TestGorm(t *testing.T) {

	// 初始化配置文件
	core.InitConf()
	// 初始化日志
	core.InitLogrus()
	defer core.CloseLogFile()
	// 初始化Gorm
	core.InitGorm()
	defer core.CloseGormLogFile()

	getOrderByOrderNum()

}

func createOrder() {
	orderItem1 := model.OrderItem{
		OrderNum:  "88888",
		ProductID: 505,
		Quantity:  2,
		UnitPrice: 20,
	}
	orderItem2 := model.OrderItem{
		OrderNum:  "88888",
		ProductID: 505,
		Quantity:  2,
		UnitPrice: 20,
	}

	var orderItems []model.OrderItem
	orderItems = append(orderItems, orderItem1, orderItem2)

	order := model.Order{
		CreatedAt:   time.Now(),
		OrderNum:    "8888",
		UserID:      2112,
		OrderStatus: 1,
		PayStatus:   1,
		PayMethod:   1,
		PayAt:       time.Now(),
		TotalAmount: 88,
		CompletedAt: time.Now(),
		CancelledAt: time.Now(),
		Note:        "test",
		OrderItems:  orderItems,
	}

	err := global.DB.Create(&order).Error
	if err != nil {
		global.Logrus.Error(err)
		return
	}
	global.Logrus.Info("添加成功")
}

func getOrderByOrderNum() {
	var order model.Order
	order.OrderNum = "8888"
	err := global.DB.Preload("OrderItems").Find(&order).Error
	if err != nil {
		global.Logrus.Error(err)
		return
	}
	fmt.Println(order)

}
