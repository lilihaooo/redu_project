package craw

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"redu/model"

	"net/url"
	"redu/flag/craw/common"
	"redu/flag/craw/receive"
	"redu/global"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 解析url
func urlInit2(page int, categoryID string, tagID uint) string {
	// 获取当前时间
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	// by=bottom_price  sort=ASC

	urlStr := "https://e.reduxingxuan.com/open/offical/product/index-search-v2?" +
		"kw=&cm=&pp0=&pp1=&f0=&f1=&pf=&" +
		"by=id&" +
		"xxbp=0&xxgy=0&jrsx=&" +
		"xsjl=&" +
		"sort=&" +
		"platform=&" +
		"type=&" +
		"page=" + strconv.FormatInt(int64(page), 10) + "&" +
		"member_id=&" +
		"product_category_id=&" +
		"product_category_two_id=&" +
		"product_category_three_id=" + categoryID + "&" +
		"is_guarantee=&" +
		//"product_tag_ids=" + strconv.Itoa(int(tagID)) + "&" +
		"product_tag_ids=&" +
		"is_displayreward=1&" +
		"is_newyeartd=&" +
		"timestamp=" + strconv.FormatInt(timestamp, 10) + "&sign_type=md5"
	// 解析URL
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("解析URL时出错:", err)
		return ""
	}
	// 获取查询参数部分
	queryParams := u.Query()
	// 删除名为"sign"的参数
	delete(queryParams, "sign")
	// 将参数按照字母顺序排序
	keys := make([]string, 0, len(queryParams))
	for key := range queryParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	// 构建排序后的查询参数字符串
	var sortedParams []string
	for _, key := range keys {
		value := queryParams[key]
		for _, v := range value {
			sortedParams = append(sortedParams, fmt.Sprintf("%s=%s", key, v))
		}
	}

	a := strings.Join(sortedParams, "&")

	// 构建新的查询参数字符串
	newQueryString := strings.Join(sortedParams, "&") + "&secret=EC7259DFCAD1715C23D400B7A88407A3"
	hash := md5.Sum([]byte(newQueryString))
	// 将哈希值转换为十六进制格式的字符串，并转换为大写形式
	hashStr := strings.ToUpper(hex.EncodeToString(hash[:]))

	u.RawQuery = a + "&sign=" + hashStr

	// 登陆
	// &member_id=214470
	//url := u.String()+"&access-token=CFP0sh5mftuB-2zCuJrGkl-FkIex7mpn"
	return u.String()
	//endUrl := u.String()+"&access-token=CFP0sh5mftuB-2zCuJrGkl-FkIex7mpn()

}

// 发请求, 抓取数据 返回data json
func grab2(wg *sync.WaitGroup, categoryID string, tagID uint) {
	// 直接抓取0到1000页, 抓取失败终止循环
	for i := 0; i <= 500; i++ {
		urlStr := urlInit2(i, categoryID, tagID)
		jsonData := ""
		for j := 0; j < 5; j++ {
			jsonData = common.SendHttp(urlStr)
			if jsonData != "" {
				break
			}
		}
		if jsonData == "[]" {
			fmt.Println("没有数据了")
			break
		}
		// 将data转为结构体
		var productRes []receive.ProductReceive
		// 使用json.Unmarshal函数解析JSON字符串到结构体变量中
		err := json.Unmarshal([]byte(jsonData), &productRes)
		if err != nil {
			fmt.Println("第"+strconv.Itoa(i)+"页, 解析JSON失败:", err)
			break
		}

		// 将接收结构体切片, 转为model结构体切片
		var products []model.Product
		for _, p := range productRes {
			product := p.ProductReceive2model()
			products = append(products, product)
		}

		var productTag model.ProductTag
		for _, product := range products {
			// 先将商品tag保存
			productTag.ProductID = product.ID
			productTag.TagID = tagID
			// 先判断该数据是否存在
			err = global.DB.Where("product_id = ? And tag_id = ?", product.ID, tagID).Take(&productTag).Error
			if err == nil {
				fmt.Println("商品的标签已经存在")
				continue
			}
			err = global.DB.Create(&productTag).Error
			if err != nil {
				fmt.Println("商品id为:" + strconv.Itoa(int(product.ID)) + " 的商品添加失败: " + err.Error())
				continue
			}
			product.ProductCategoryId = categoryID
			// 商品存在, 由于主键唯一都不会被插入
			err = global.DB.Create(&product).Error
			if err != nil {
				fmt.Println("商品id为:" + strconv.Itoa(int(product.ID)) + " 的商品添加失败: " + err.Error())
				continue
			}
			fmt.Println("ok")
		}
	}
	defer wg.Done()
}

func Task2(tagID uint) {
	// 查询所有三级分类的分类id
	var c model.Category
	var categoryIDs []string
	global.DB.Model(c).Where("level = 3").Select("id").Scan(&categoryIDs)

	var wg sync.WaitGroup
	num := len(categoryIDs)
	wg.Add(num)
	for _, categoryID := range categoryIDs {
		go grab2(&wg, categoryID, tagID)
	}
	wg.Wait()
	var count int64
	global.DB.Model(model.Product{}).Count(&count)
	fmt.Printf("一共" + strconv.FormatInt(count, 10) + "条商品记录")
}
