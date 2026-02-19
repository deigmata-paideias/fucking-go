package main

import (
	"fmt"
	"math/big"
)

// 经典语言问题
// float 类型的 a + b 不等于 a + b

func main() {

	var a, b = 0.1, 0.2
	var c = a + b
	var d = 0.3

	fmt.Println(c == d)
	fmt.Println("a + b =", c)

	// 上述程序的输出结果是：
	// Macos M4 芯片
	// false
	// a + b = 0.30000000000000004
	// 可以看出是有差异的

	fmt.Println("===================================")

	// 其中一种解决方法是：使用 math/big 包中的 big.Float 类型来进行高精度的浮点数计算。
	var e, f = big.NewFloat(0.1), big.NewFloat(0.2)
	var g = new(big.Float).Add(e, f)
	var h = big.NewFloat(0.3)

	// 在判断相等时，使用一个差值来比较
	fmt.Println(EqualFloat(g, h))
	fmt.Println("e + f = ", g)
}

// CustomFloat 定义一个类型约束，支持 float64 和 float32
type CustomFloat interface {
	float64 | float32 | *big.Float
}

// float 类型比较的容差值
const epsilon = 1e-10

func EqualFloat[T CustomFloat](a, b T) bool {

	// 使用 *big.Float 进行比较
	var bigA, bigB *big.Float

	switch v := any(a).(type) {
	case float64:
		bigA = big.NewFloat(v)
	case float32:
		bigA = big.NewFloat(float64(v))
	case *big.Float:
		bigA = v
	}

	switch v := any(b).(type) {
	case float64:
		bigB = big.NewFloat(v)
	case float32:
		bigB = big.NewFloat(float64(v))
	case *big.Float:
		bigB = v
	}

	diff := new(big.Float).Sub(bigA, bigB)
	return diff.Abs(diff).Cmp(big.NewFloat(epsilon)) < 0
}
