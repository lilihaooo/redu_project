package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	// 测试 Add 函数的正确性
	str := urlDecrypt(1, 2)

	fmt.Println(str)

}

// Add 函数接受两个整数并返回它们的和

func urlDecrypt(page int, categoryID int64) string {
	// 获取当前时间
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	// by=bottom_price  sort=ASC

	urlStr := "https://e.reduxingxuan.com/open/offical/product/index-search-v2?" +
		"kw=&cm=&pp0=&pp1=&f0=&f1=&pf=&" +
		"by=&xxbp=0&xxgy=0&jrsx=0&xsjl=0&sort=&" +
		"platform=&" +
		"type=&" +
		"page=" + strconv.FormatInt(int64(page), 10) + "&" +
		"member_id=&" +
		"product_category_id=&" +
		"product_category_two_id=&" +
		"product_category_three_id=42&" +
		"is_guarantee=0&" +
		"product_tag_ids=&" +
		"is_displayreward=&" +
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
}
