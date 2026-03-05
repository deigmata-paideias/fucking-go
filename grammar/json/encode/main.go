package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	// go 的编码器默认会加一个换行符
	// JSON编码器的设计是用于流处理的。JSON的流式处理通常意味着使用一个换行符分隔多个JSON对象，这就是为什么JSON编码器会在最后添加一个换行符的原因
	data := map[string]int{"key": 1}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)

	raw, _ := json.Marshal(data)

	if b.String() == string(raw) {
		fmt.Println("same encoded data")
	} else {
		fmt.Printf("'%s' != '%s'\n", raw, b.String())
		//prints:
		//'{"key":1}' != '{"key":1}\n'
	}
}
