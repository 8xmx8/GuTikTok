package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//=== RUN   TestConsulService
//http://127.0.0.1:8500
//Status: 200
//--- PASS: TestConsulService (0.00s)
//PASS

func TestConsulService(t *testing.T) {
	// 创建一个模拟的 Consul 服务
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// 在测试完成后关闭模拟的 Consul 服务
	defer server.Close()

	// 获取模拟 Consul 服务的地址
	consulURL := "http://127.0.0.1:8500"

	fmt.Println(consulURL)
	// 发送 GET 请求到模拟的 Consul 服务
	resp, err := http.Get(consulURL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 获取响应的状态码
	statusCode := resp.StatusCode

	// 判断请求是否成功
	if statusCode != http.StatusOK {
		t.Errorf("Request failed. Status Code: %d", statusCode)
	}
	fmt.Println("Status:", statusCode)
}
