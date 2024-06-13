package flag

import (
	"flag"
	"fmt"
	"github.com/fatih/structs"
	"redu/flag/craw"
	"redu/flag/db"
	"redu/flag/fake"
	"time"
)

type Option struct {
	DB   bool
	Craw string
	Fake string
}

// Parse 解析命令行参数
func Parse() Option {
	db_ := flag.Bool("db", false, "初始化数据库") // false 默认值, 含有-db就为true
	fake_ := flag.String("fake", "", "生成模拟数据")
	craw_ := flag.String("craw", "", "启动爬虫")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	return Option{
		DB:   *db_,
		Craw: *craw_,
		Fake: *fake_,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	crMap := structs.Map(&option)
	for _, v := range crMap {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}

	}
	return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		db.Migrate()
		return
	}
	// 伪造数据

	if option.Fake == "MES" { // 生成体验分
		fake.Task()
		return
	}
	if option.Fake == "MP" { // 生成体验分
		a := time.Now()
		fake.Task2()
		b := time.Since(a)
		fmt.Println(b)
		return
	}

	// 爬虫
	if option.Craw == "product" {
		craw.Task()
		return
	}
	if option.Craw == "product_tag" {
		craw.Task2(1)
		return
	}
	if option.Craw == "category" {
		craw.CategoryTask()
		return
	}

}
