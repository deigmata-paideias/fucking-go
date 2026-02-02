package main

import "os"

func main() {

	f, err := os.OpenFile(
		"a.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		panic(err)
	}

	// close 函数有返回值，返回值是一个 err
	// defer f.Close()

	// 检查下 err，可以确保程序健壮性。
	// 为此可以写出来 V1 的检查版本，check err
	// 	这种写法会 panic 吗？答案是：会 panic
	//  => panic: write a.txt: file already closed
	//
	// 在 Go 中，每次"defer"语句执行时，函数值和调用的参数都会照常求值并重新保存，但实际函数不会被调用
	// 	也就是说，checkErrV1 里的函数被执行了，f.Close() 也被执行了
	// 	但是 f.Close() 的返回值没有被任何人接收和处理
	// 	这就导致文件被关闭了，后续的写操作就会 panic
	// 参考：https: //golang.ac.cn/ref/spec#Defer_statements
	//
	//defer checkErrV1(f.Close())

	// 因此正确的检查写法是
	defer checkErrV2(f.Close)

	_, err = f.WriteString("append, hello")
	if err != nil {
		panic(err)
	}
}

func checkErrV1(err error) {

	if err != nil {
		panic(err)
	}
}

func checkErrV2(f func() error) {

	if err := f(); err != nil {
		panic(err)
	}
}
