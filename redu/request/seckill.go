package request

import (
	"redu/serialize/common"
)

type CreateSeckillActivityParam struct {
	ProductID      uint                 `json:"product_id"`
	Title          string               `json:"title" validate:"required" label:"标题"`
	Status         common.SeckillStatus `json:"status" validate:"oneof=0 1 2" label:"状态"`
	Count          int                  `json:"count" validate:"required" label:"数量"`
	UnitPrice      float64              `json:"unit_price" validate:"required" label:"单价"`
	BeginTime      string               `json:"begin_time" validate:"required" label:"开始时间"`
	EndTime        string               `json:"end_time" validate:"required" label:"结束时间"`
	AllowUserGrade int                  `json:"allow_user_grade" validate:"required" label:"允许的用户等级"`
}

type GetSeckillListParam struct {
	PaginationParam
}

type UpdateSeckillActivityParam struct {
	ID             uint    `json:"id"`
	Count          int     `json:"count" validate:"required" label:"数量"`
	UnitPrice      float64 `json:"unit_price" validate:"required" label:"单价"`
	BeginTime      string  `json:"begin_time" validate:"required" label:"开始时间"`
	EndTime        string  `json:"end_time" validate:"required" label:"结束时间"`
	AllowUserGrade int     `json:"allow_user_grade" validate:"required" label:"允许的用户等级"`
}
