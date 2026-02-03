package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var val = struct {
		NoteBook any
		Flag     any `json:"notEBook"`
	}{"", false}

	// 下面的代码会输出什么？
	//  	答案是 {true false}
	// 为什么？预期中应该是 A 和 true。
	//
	// 这和 Go json 包的 tag 解析 json 规则有关。
	//
	// 参考：https://pkg.go.dev/encoding/json#Unmarshal
	//   To unmarshal JSON into a struct, Unmarshal matches incoming object keys to the keys used by Marshal (either the struct field name or its tag), ignoring case. If multiple struct fields match an object key, an exact case match is preferred over a case-insensitive one.
	//   翻译过来就是：要将 JSON 反序列化为结构体，`Unmarshal` 会将传入对象的键与 `Marshal` 使用的键（结构体字段名或其标签）进行匹配，不区分大小写。如果多个结构体字段与对象键匹配，则精确的大小写匹配优先于不区分大小写的匹配。
	//
	// 但是不太能解释这里的现象，NoteBook 和 NoteBook 是符合精确大小写匹配规则的？但是 A 没有被写进去
	//
	// 在 go.dev/blog/json 中有以下说明：
	// How does Unmarshal identify the fields in which to store the decoded data? For a given JSON key "Foo", Unmarshal will look through the destination struct’s fields to find (in order of preference):
	//
	//  An exported field with a tag of "Foo" (see the Go spec for more on struct tags),
	//
	//  An exported field named "Foo", or
	//
	//  An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".
	// 翻译是：
	//    1. 按照大小写敏感的方式匹配 tag；
	//    2. 按照大小写敏感的方式匹配字段名；
	//    3. 按照大小写不敏感的方式匹配字段名或者 tag。
	//
	// 结合上面的描述，可以推测出：输出的结果是符合预期的。
	//
	// 结论就是：自己项目中定义明确的 json tag，都用蛇形 tag 定义，按照规则匹配（但是第三方系统的 json 格式比较难要求
	err := json.Unmarshal([]byte(`{"NoteBook": "A", "notEbook":"true"}`), &val)
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
