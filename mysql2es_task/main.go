package main

import (
	"mysql2es_task/core"
	"mysql2es_task/service/product_ser"
)

func main() {
	core.InitConf()
	core.InitEs()
	core.InitGorm()
	defer core.CloseGormLogFile()
	core.InitLogrus()
	defer core.CloseLogFile()

	product_ser.Mysql2Es()
}
