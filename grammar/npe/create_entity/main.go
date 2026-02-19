package main

import "fmt"

type User struct {
	ID   uint
	Name string
	Age  int8
}

func main() {

	// 在此中场景下，变量是一个指向 User 结构题的指针，但没有初始化。访问字段会出现 NPE。
	// 如果使用 Goland，在 user1.Name 会提示：Potential nil dereference
	//
	// 运行时产生：
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x10 pc=0x100325474]
	var user1 *User
	user1.Name = "Alice"
	fmt.Println(user1)

	// 通过 new 函数创建一个 User 结构体的实例，并返回一个指向该实例的指针。此时，user2 已经被初始化，可以安全地访问其字段。
	user2 := new(User)
	user2.Name = "Bob"
	fmt.Println(user2)

	// 直接使用 &User{} 创建一个 User 结构体的实例，并返回一个指向该实例的指针。此时，user3 已经被初始化，可以安全地访问其字段。
	user3 := &User{}
	user3.Name = "Charlie"
	fmt.Println(user3)

	// 直接使用 User{} 创建一个 User 结构体的实例，并返回一个值。此时，user4 是一个 User 结构体的值，可以直接访问其字段。
	user4 := User{}
	user4.Name = "David"
	fmt.Println(user4)

}
