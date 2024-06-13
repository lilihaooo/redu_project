package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendHttp(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求时出错:", err)
		return ""
	}
	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	// 发送 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("发送请求出错:", err)
		return ""
	}
	defer resp.Body.Close()
	reader := resp.Body
	// 读取响应体数据
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("读取响应体出错:", err)
		return ""
	}
	return string(body)
}
