package main

import (
	"log"

	"github.com/deigmata-paideias/fucking-go/gorm"
	pkggorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 踩坑点：
// Gorm 对于 User 结构体，如果没有重写 TableName 方法，默认会将结构体名转换为 snake_case 的复数形式作为表名，即 "users"。
//
// 从下面的 SQL 日志中可以看到，数据库表名自动加了 s。
//
// 2026/02/15 22:37:41 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/db.go:18
// [0.043ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
//
// 2026/02/15 22:37:41 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/db.go:18
// [1.716ms] [rows:0] CREATE TABLE `users` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text,`age` integer)
//
// 2026/02/15 22:37:41 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/table_name/main.go:30
// [0.747ms] [rows:1] INSERT INTO `users` (`name`,`age`) VALUES ("Alice",30) RETURNING `id`
//
// 2026/02/15 22:39:59 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/db.go:18
// [0.035ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type='table' AND name="users"
//
// 2026/02/15 22:39:59 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/db.go:18
// [1.604ms] [rows:0] CREATE TABLE `users` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text,`age` integer)
//
// 2026/02/15 22:39:59 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/db.go:18
// [0.008ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type='table' AND name="data_resources"
//
// 2026/02/15 22:39:59 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/db.go:18
// [0.520ms] [rows:0] CREATE TABLE `data_resources` (`id` integer PRIMARY KEY AUTOINCREMENT,`name` text,`desc` text)
//
// 2026/02/15 22:39:59 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/table_name/main.go:45
// [0.551ms] [rows:1] INSERT INTO `users` (`name`,`age`) VALUES ("Alice",30) RETURNING `id`
//
// 2026/02/15 22:39:59 /Users/shown/workspace/golang/open_source/fucking-go/project/gorm/table_name/main.go:50
// [0.510ms] [rows:1] INSERT INTO `data_resources` (`name`,`desc`) VALUES ("Alice-Resource","This is Alice's resource") RETURNING `id`

type User struct {
	ID   uint
	Name string
	Age  int8
}

type DataResource struct {
	ID   uint
	Name string
	Desc string
}

var db *pkggorm.DB

func init() {

	_db := gorm.Init("./table_name/gorm.db", []any{&User{}, &DataResource{}}, &pkggorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if _db == nil {
		log.Fatal("Init db failed!")
	}

	db = _db
}

func main() {

	db.Create(&User{
		Name: "Alice",
		Age:  30,
	})

	db.Create(&DataResource{
		Name: "Alice-Resource",
		Desc: "This is Alice's resource",
	})
}
