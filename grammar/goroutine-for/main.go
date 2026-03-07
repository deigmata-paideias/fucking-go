package main

import (
	"fmt"
	"time"
)

func main() {

	// 下面的代码在更早之前（go 1.22）输出的一直是 e，对应的解决方法是在 go 后面函数声明一个变量传递
	// 因为 range 循环在这之前是复用 v 变量的，1.22 之后改为迭代变量，没有此问题
	data := []string{"a", "b", "c", "d", "e"}

	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(time.Second)
}
