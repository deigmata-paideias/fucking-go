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

	// go 中访问 map 不存在的 key 不会报错，返回值为类型的零值
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	if v := x["two"]; v == "" { //incorrect
		fmt.Println("no entry")
	}

	// 正确写法
	if _, ok := x["two"]; !ok {
		fmt.Println("no entry")
	}

	// 字符串变量是不可变的
	s := "hello"
	// s[0] = "H"
	// invalid operation: s[0] (cannot assign to s[0])
	fmt.Println(s)

	// 如果需要改变，需要转为 byte 在修改
	s = "hello"
	sByte := []byte(s)
	sByte[0] = 'H'
	s = string(sByte)
	fmt.Println(s)

}
