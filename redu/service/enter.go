package service

import (
	"redu/service/product"
	"redu/service/seckill"
)

type Service struct {
	product.Product
	seckill.Seckill
}

var AppService = new(Service)
