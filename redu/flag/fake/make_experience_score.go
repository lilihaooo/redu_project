package fake

import (
	"fmt"
	"math"
	"math/rand"
	"redu/global"
	"redu/model"
	"strconv"
	"sync"
	"time"
)

var totalCount int
var mutex sync.Mutex

func makeExperienceScore(wg *sync.WaitGroup, offset int) {
	defer wg.Done()
	fmt.Println("开始")
	pageSize := 10000
	var product model.Product
	var productIDs []uint
	err := global.DB.Model(product).Select("id").Order("id desc").Limit(pageSize).Offset(offset).Scan(&productIDs).Error
	if err != nil {
		global.Logrus.Error(err)
	}

	for _, productID := range productIDs {
		err = global.DB.Model(product).Where("id = ?", productID).Update("experience_score", getScore()).Error
		if err != nil {
			fmt.Println("ID为: " + strconv.Itoa(int(product.ID)) + " 的商品修改失败")
		} else {
			mutex.Lock()
			totalCount++
			mutex.Unlock()
		}
		count := totalCount
		if count%100 == 0 {
			fmt.Printf("总共已执行 %d 次\n", count)
		}

	}

}

func getScore() float64 {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 生成1到50之间的随机整数
	randomInt := rand.Intn(50) + 1
	// 转换为浮点数并除以10，得到保留一位小数的结果
	return float64(randomInt) / 10

}

func Task() {
	start := time.Now()
	var wg sync.WaitGroup
	var total int64
	var model model.Product
	global.DB.Model(model).Count(&total)
	limit := 10000
	manNum := int(math.Ceil(float64(total) / float64(limit)))
	wg.Add(manNum)
	for i := 1; i <= manNum; i++ {
		offset := (i - 1) * 10000
		go makeExperienceScore(&wg, offset)
	}
	wg.Wait()
	// 计算函数执行的时间
	duration := time.Since(start)

	//打印函数执行的时间
	fmt.Println("完成耗时: ", duration)
	fmt.Println("总记录: " + strconv.Itoa(totalCount))
}
