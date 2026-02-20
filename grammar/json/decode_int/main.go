package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 这个例子是为了演示 Go 语言中 JSON 解码时，整数类型的特殊处理。Go 的 JSON 解码器在处理整数时，会将其解码为 float64 类型，以避免整数溢出的问题。
// 这可能会导致一些意想不到的问题，特别是在处理大整数时。

func main() {

	// 需要解析的 json 数据
	var data = []byte(`{"status": 200}`)

	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	// 输出类型，预计中是 int64？但是实际是 float64
	fmt.Printf("%T\n", result["status"])

	// 产生类型断言错误
	// panic: interface conversion: interface {} is float64, not int
	var status = result["status"].(int)
	fmt.Println("Status value: ", status)

	// 解决办法有四种，列举下简单的两种
	// 1. 类型转换，明确知道自断类型，在 decode 之后强制转换下
	// 紧接着上面的 result
	result["status"] = int(result["status"].(float64))
	fmt.Printf("%T\n", result["status"])
	fmt.Println("Status value: ", result["status"])

	// 2. 使用结构体 tag 序列化
	type Data struct {
		Status int `json:"status"`
	}
	var structData *Data
	err := json.Unmarshal(data, &structData)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%T\n", structData.Status)
	fmt.Println("Status value: ", structData.Status)
}
