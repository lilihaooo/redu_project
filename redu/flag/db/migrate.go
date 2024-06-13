package db

import (
	"redu/global"
	"redu/model"
)

func Migrate() {
	var err error
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.Category{},
		)
	if err != nil {
		global.Logrus.Errorf("生成数据库表结构失败:%s", err)
		return
	}
	global.Logrus.Info("生成数据库表结构成功")
}
