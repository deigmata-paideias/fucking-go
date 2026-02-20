package main

import "fmt"

// recover 只有在 defer 执行的函数中才能生效

// 错误使用示例
//func main() {
//
//	// 在 panic 之前，无任何作用
//	recover()
//	panic("panic demo")
//
//	// Goland 提示 Unreachable code
//	recover()
//}

func main() {

	defer func() {
		fmt.Println("recover exec:", recover())
	}()

	panic("panic demo")
}
