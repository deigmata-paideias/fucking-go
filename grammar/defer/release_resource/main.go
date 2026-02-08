package main

import "fmt"

func main() {

	// 申请一个 R1 资源
	r := NewResource("R1")
	// 使用 defer 及时释放
	defer r.Release()

	fmt.Println("资源名：", r.Name)

	// 申请一个 R2 资源，但是使用 r 变量
	// 重新赋值之前，确保 r 已经释放了，手动调用下
	r.Release()
	r = NewResource("R2")

	fmt.Println("资源名：", r.Name)

	// 运行之后发现输出是：
	//	资源名： R1
	//	R1释放了
	//	资源名： R2
	//	R1释放了
	//
	// 这里的问题是：R1 被释放了两次，第一次是 defer 里，第二次是手动调用的 r.Release()，这就导致了资源被重复释放了
	// 而 R2 没有被释放，会造成资源泄漏。
	// 这里的问题和 check-error 一样，因为 defer 后面的函数或者参数是被保存起来的
	// 并没有现场执行，而是延迟在函数退出时候执行，
	// 解决方法也是用匿名函数，闭包封装下，像下面这样
	// 关键原理是闭包函数中的 r 是共享的
	// defer func() {
	//  	r.Release()
	// }
}

type Resource struct {
	Name string
}

func NewResource(name string) *Resource {

	return &Resource{Name: name}
}

func (r *Resource) Release() {

	fmt.Println(r.Name + "释放了")
}
