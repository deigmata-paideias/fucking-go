package main

import "fmt"

func main() {

	data := []any{1, 2, 3, 4, 5}

	// 可变参数
	fmt.Println(data)
	fmt.Println(data...)

	// map 遍历不是有序的
	dataMap := map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("第 %d 次输出: ", i+1)
		for k, v := range dataMap {
			fmt.Printf("%s: %v \t", k, v)
		}
		fmt.Println()
	}
}
