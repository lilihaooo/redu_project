package v1

import (
	"redu/api/v1/product"
	"redu/api/v1/seckill"
)

type Group struct {
	product.Product
	seckill.Seckill
}

var GroupApp = new(Group)
