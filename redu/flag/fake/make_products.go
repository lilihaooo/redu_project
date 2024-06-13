package fake

import (
	"fmt"
	"math"
	"math/rand"
	"redu/global"
	"redu/model"
	"strconv"
	"time"
)

func makeProduct() {
	var total int64

	global.DB.Table("product").Count(&total)

	for {
		a := time.Now()
		if total > 10000000 {
			break
		}
		// 初始化随机数生成器
		rand.Seed(time.Now().UnixNano())
		offset := rand.Intn(int(total)) + 1 - 300
		if offset < 0 {
			offset = 0
		}

		products := getProducts(offset)

		newP := changeValue(products)

		db := global.DB.Create(&newP)
		err := db.Error
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("受影响行: " + strconv.FormatInt(db.RowsAffected, 10))
		fmt.Println("总记录: " + strconv.FormatInt(total, 10))
		b := time.Since(a)
		fmt.Println(b)
		total += 1000
	}
}

func getProducts(offset int) (products []model.Product) {
	var IDs []int64
	sql := "SELECT Id FROM product  ORDER BY id asc LIMIT " + strconv.Itoa(offset) + ", 100"
	err := global.DB.Raw(sql).Scan(&IDs).Error

	if err != nil {
		global.Logrus.Fatal(err)
	}
	global.DB.Where("id in ?", IDs).Preload("ProductImages").Find(&products)
	return
}

func Task2() {
	makeProduct()
}

func changeValue(products []model.Product) (newP []model.Product) {
	for i := 1; i <= 10; i++ {
		for _, one := range products {
			var product model.Product
			var imgs []*model.ProductImage
			img := model.ProductImage{
				Url: one.ProductImages[0].Url,
			}
			imgs = append(imgs, &img)
			product.ProductImages = imgs
			product.ShopID = one.ShopID
			//bdate := getRandomDateIn2023()
			//edate := addRandomDaysToDate(bdate, 100)
			//product.BeginTime = int(bdate.Unix())
			//product.EndTime = int(edate.Unix())
			product.XingxuanCredit = getXingxuanCredit()
			product.SalesAll = strconv.Itoa(get0To10000())
			product.Sales24 = strconv.Itoa(get0To10000())
			product.Platform = one.Platform
			product.Title = one.Title
			cm := get0To80()
			product.Commission = strconv.Itoa(cm)
			product.BottomPrice = round(getBottomPrice(), 2)
			//product.Url = one.Url
			product.ProductCategoryId = one.ProductCategoryId
			//product.SpecimenCount = one.SpecimenCount
			//product.RewardBeginTime = one.RewardBeginTime
			//product.RewardEndTime = one.RewardEndTime
			product.TrueSales = strconv.Itoa(get0To10000())
			cmv := product.BottomPrice * float64(cm) * 0.01
			product.CommissionValue = round(cmv, 2)
			product.Remaining = get0To1000000()
			product.ExperienceScore = getScore()
			product.Status = "1"
			product.IsFake = 1
			newP = append(newP, product)
		}
	}
	return
}

func getRandomDateIn2023() time.Time {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
	// 随机月份（1-12）
	month := rand.Intn(12) + 1
	// 根据月份计算随机日期
	var day int
	switch month {
	case 2: // 二月可能是28天或29天（闰年）
		if (2023%4 == 0 && 2023%100 != 0) || 2023%400 == 0 {
			day = rand.Intn(29) + 1
		} else {
			day = rand.Intn(28) + 1
		}
	case 4, 6, 9, 11: // 30天的月份
		day = rand.Intn(30) + 1
	default: // 其他月份（1, 3, 5, 7, 8, 10, 12）是31天
		day = rand.Intn(31) + 1
	}
	// 构建随机日期
	randomDate := time.Date(2023, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	// 返回随机日期
	return randomDate
}

func addRandomDaysToDate(date time.Time, maxDays int) time.Time {
	// 初始化随机数生成器（如果需要，这里也可以使用全局的随机数生成器）
	rand.Seed(time.Now().UnixNano())
	// 随机添加1到maxDays之间的天数
	daysToAdd := rand.Intn(maxDays) + 1
	return date.AddDate(0, 0, daysToAdd)
}

func getXingxuanCredit() string {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 定义元素切片
	elements := []string{"A级", "B级", "S级"}

	// 随机选择一个元素
	randomIndex := rand.Intn(len(elements))
	return elements[randomIndex]
}

func get0To10000() int {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// rand.Intn(n) 生成一个[0, n)之间的随机整数
	// 因此，要生成1到10000之间的随机数，需要加1
	return rand.Intn(10000) + 1
}

func get0To1000000() int {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// rand.Intn(n) 生成一个[0, n)之间的随机整数
	// 因此，要生成1到10000之间的随机数，需要加1
	return rand.Intn(10000) + 1
}

func get0To80() int {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// rand.Intn(n) 生成一个[0, n)之间的随机整数
	// 因此，要生成1到10000之间的随机数，需要加1
	return rand.Intn(80) + 1
}

func getBottomPrice() float64 {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 生成一个[0, 2000000)之间的随机整数
	// 因为0.01 * 2000000 = 20000，所以乘以0.01后会落在[0.01, 20000)之间
	// 最后我们再加0.01确保最大值是20000
	randomInt := rand.Intn(2000000)
	randomFloat := float64(randomInt) * 0.01
	if randomFloat == 20000 {
		// 如果生成的随机数恰好是20000，则直接返回，因为边界值包含在内
		return randomFloat
	}
	// 如果不是20000，则加0.01确保最大值是20000
	return math.Min(randomFloat+0.01, 20000)
}

// round 函数将浮点数四舍五入到指定的小数位数
func round(value float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Floor(value*shift+0.5) / shift
}
