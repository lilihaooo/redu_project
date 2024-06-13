package v2

import (
	"redu/api/v2/product"
)

type Group struct {
	product.Product
}

var GroupApp = new(Group)
