package model

import "redu/global"

type Category struct {
	ID       string     `json:"id"`
	Level    string     `json:"level"`
	ParentID string     `json:"parent_id"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Image    string     `json:"image"`
	Num      int        `json:"num"`
	Sort     int        `json:"sort"`
	Childer  []Category `json:"childer" gorm:"foreignkey:ParentID"`
}

func (c *Category) GetAll() (list []Category, err error) {
	err = global.DB.Order("level asc, sort asc").Find(&list).Error
	return
}
