package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	for count := 0; count < 10; count++ {
		getSignleNumber()
	}
}
func getSignleNumber() {
	// 基础API地址
	baseURL := "https://api.sms-activate.ae/stubs/handler_api.php"

	// 创建查询参数
	params := url.Values{}
	params.Set("api_key", "c6de5132c013098Ad74b4d4dc6b4b609") // 替换为你的实际API密钥
	params.Set("action", "getNumberV2")
	params.Set("service", "tg")
	params.Set("forward", "0")
	params.Set("operator", "") // 空值参数保留
	params.Set("ref", "")      // 空值参数保留
	params.Set("country", "0")
	params.Set("phoneException", "")
	params.Set("maxPrice", "")
	params.Set("useCashBack", "")
	params.Set("activationType", "0")
	params.Set("language", "en-US")
	params.Set("userId", "") // 空值参数保留

	// 构造完整URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建请求对象
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return
	}

	// 设置请求头（可选）
	req.Header.Set("User-Agent", "SMSActivateClient/1.0 (Go)")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 输出结果
	fmt.Printf("HTTP状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容:\n%s\n", string(body))
}
