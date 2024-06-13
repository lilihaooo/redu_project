package fake

import (
	"fmt"
	"redu/global"
	"redu/model"
	"strconv"
)

func Task3() {
	var ids []int64
	err := global.DB.Model(model.Product{}).Limit(50000).Where("title like ?", "Fake%").Select("id").Scan(&ids).Error
	if err != nil {
		global.Logrus.Error(err)
		return
	}

	tx := global.DB.Begin()

	err = tx.Where("id in ?", ids).Delete(model.Product{}).Error
	if err != nil {
		tx.Rollback()
		global.Logrus.Error(err)
		return
	}
	err = tx.Where("product_id in ?", ids).Delete(model.ProductImage{}).Error
	if err != nil {
		tx.Rollback()
		global.Logrus.Error(err)
		return
	}
	tx.Commit()
	fmt.Println("删除" + strconv.Itoa(len(ids)) + "条数据")
}
