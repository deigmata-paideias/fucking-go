package main

import (
	"errors"
	"fmt"
	"log"

	pkggorm "gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/deigmata-paideias/fucking-go/gorm"
)

// Gorm 在什么情况下会返回 ErrRecordNotFound 错误？
// 在之前的 Gorm 固有印象中，凡是没有找到 record 的，都会返回 ErrRecordNotFound

var db *pkggorm.DB

func init() {

	_db := gorm.Init("./err_record_not_found/gorm.db", []any{&User{}}, &pkggorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if _db == nil {
		log.Fatal("Init db failed!")
	}

	db = _db
}

type User struct {
	ID   uint
	Name string
	Age  int8
}

func main() {

	// First 方法查询
	firstResUser := &User{}
	firstResUserSlice := make([]User, 0)
	err1 := db.Model(firstResUser).Where("id = ?", 1).First(firstResUser).Error
	if errors.Is(err1, pkggorm.ErrRecordNotFound) {
		log.Println("First 方法使用结构体接受返回值，没有找到记录，返回了 ErrRecordNotFound 错误")
	} else {
		log.Println("First 方法使用结构体接受返回值，没有找到记录，返回了其他错误或未返回错误 ")
		log.Printf("Err: %v\n", err1.Error())
	}
	fmt.Printf("\n%+v\n", firstResUser)

	err2 := db.Model(firstResUser).Where("id = ?", 1).First(&firstResUserSlice).Error
	if errors.Is(err2, pkggorm.ErrRecordNotFound) {
		log.Println("First 方法使用 Slice 接受返回值，没有找到记录，返回了 ErrRecordNotFound 错误")
	} else {
		log.Println("First 方法使用结构体接受返回值，没有找到记录，返回了其他错误或未返回错误 ")
		log.Printf("Err: %v\t, res slice len: %v", err2, len(firstResUserSlice))
	}
	fmt.Printf("\n%+v\n", firstResUserSlice)

	fmt.Println("====================================")

	// Find 方法查询
	findResUser := &User{}
	findResUserSlice := make([]User, 0)
	err3 := db.Model(findResUser).Where("id = ?", 1).Find(findResUser).Error
	if errors.Is(err3, pkggorm.ErrRecordNotFound) {
		log.Println("Find 方法使用结构体接受返回值，没有找到记录，返回了 ErrRecordNotFound 错误")
	} else {
		log.Println("Find 方法使用结构体接受返回值，没有找到记录，返回了其他错误或未返回错误 ")
		log.Printf("Err: %v\n", err3)
	}
	fmt.Printf("\n%+v\n", firstResUser)

	err4 := db.Model(findResUser).Where("id = ?", 1).Find(&findResUserSlice).Error
	if errors.Is(err4, pkggorm.ErrRecordNotFound) {
		log.Println("Find 方法使用 Slice 接受返回值，没有找到记录，返回了 ErrRecordNotFound 错误")
	} else {
		log.Println("Find 方法使用结构体接受返回值，没有找到记录，返回了其他错误或未返回错误 ")
		log.Printf("Err: %v, \t res slice len: %v", err4, len(findResUserSlice))
	}
	fmt.Printf("\n%+v\n", findResUserSlice)

	// 2026/02/18 23:08:52 First 方法使用结构体接受返回值，没有找到记录，返回了 ErrRecordNotFound 错误
	// 2026/02/18 23:08:52 First 方法使用 Slice 接受返回值，没有找到记录，返回了 ErrRecordNotFound 错误
	//
	// 2026/02/18 23:08:52 Find 方法使用结构体接受返回值，没有找到记录，返回了其他错误或未返回错误
	// 2026/02/18 23:08:52 Err: <nil>
	// 2026/02/18 23:08:52 Find 方法使用结构体接受返回值，没有找到记录，返回了其他错误或未返回错误
	// 2026/02/18 23:08:52 Err: <nil>, 	 res slice len: 0
	//
	// 运行上述代码，我们会发现，在使用 Find 方法查询数据时，不论使用 Struct 或者 Slice 接受返回值，
	// Gorm 都不会返回 ErrRecordNotFound 错误，而是直接返回 nil，
	// 并且将查询结果赋值为 Struct 的零值或者 Slice 的空切片。
	// 这说明，Gorm 只有在使用 First 方法查询数据时，如果没有找到记录，
	// 才会返回 ErrRecordNotFound 错误。而在使用 Find 方法查询数据时，即使没有找到记录，
	// 也不会返回 ErrRecordNotFound 错误，而是直接返回 nil

}
