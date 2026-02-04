package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {

	type Member struct {
		ID   int
		Name string
		Age  int
	}

	members := []*Member{
		{ID: 3, Name: "Bb", Age: 30},
		{ID: 2, Name: "Aa", Age: 25},
		{ID: 1, Name: "Cc", Age: 35},
	}

	idMap := make(map[int]*Member)
	for idx := range members {
		idMap[members[idx].ID] = members[idx]
	}

	// 需求：按照 Name 字段进行排序
	// 输出 id 为 2 的字符串名是哪个？预期应该是 Aa
	// 但是实际输出是 3 Bb？
	// slices sort 函数是原地交换排序
	// slice 是一块连续的内存空间，那原地交换排序就是
	// 直接在这块内存空间上进行交换，则 map 中的指针就会指向错误的地址。
	// 解决方法是 slice 换成 []*Member 类型，直接对指针交换。或者使用 sort.Slice 函数 等其他方法
	slices.SortFunc(members, func(a, b *Member) int {
		// 按照字典序排列
		return strings.Compare(a.Name, b.Name)
	})

	fmt.Println(idMap[2])
}
