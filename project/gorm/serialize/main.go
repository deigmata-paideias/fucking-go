package main

import (
	"fmt"
	"log"

	"github.com/deigmata-paideias/fucking-go/gorm"
	pkggorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	ID uint64
	// Gorm 中的序列化器（Serializer）将复杂的数据结构（如 map、struct 等）序列化为 JSON 存储在数据库中。
	// 通过在字段标签中指定 `serializer:json`，Gorm 会自动将该字段的值序列化为 JSON 格式存储，并在查询时反序列化回原始数据结构。
	Other map[string]string `gorm:"serializer:json;type:longtext"`
	API   *API              `gorm:"serializer:json;type:longtext"`
}

type API struct {
	Name string
	Desc string
}

var db *pkggorm.DB

func init() {

	_db := gorm.Init("./table_name/gorm.db", []any{&Data{}}, &pkggorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if _db == nil {
		log.Fatal("Init db failed!")
	}

	db = _db
}

func main() {

	data := &Data{
		Other: map[string]string{
			"1": "2",
		},
		API: &API{
			Name: "api",
			Desc: "api desc",
		},
	}

	db.Create(data)

	var dbData Data
	if err := db.Model(&dbData).Where("id = ?", data.ID).First(&dbData).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("db data: %#+v", dbData)
}
