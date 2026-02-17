package main

import (
	"log"

	"github.com/deigmata-paideias/fucking-go/gorm"
	pkggorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID      uint
	Name    string
	Age     int8
	Enabled bool
}

var db *pkggorm.DB

func init() {

	_db := gorm.Init("./update_column/gorm.db", []any{&User{}}, &pkggorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if _db == nil {
		log.Fatal("Init db failed!")
	}

	db = _db
}

func main() {

	aliceUser := &User{
		Name:    "Alice",
		Age:     18,
		Enabled: true,
	}

	db.Model(&User{}).Create(aliceUser)

	// 需求描述：用户 Alice 在刚创建时处于激活状态，即 Enabled 为 true。
	// admin 超管将 Alice 设置为锁定状态，账户不可用。Enabled 字段被更新为 false。
	aliceUser.Enabled = false
	db.Model(&User{}).Updates(aliceUser)

	var dbAliceUser1 User
	db.Model(&User{}).First(&dbAliceUser1, aliceUser.ID)
	log.Printf("Alice user in db 1: %+v", dbAliceUser1)

	// 现在超管想恢复 Alice 账户，设置为 true
	aliceUser.Enabled = true
	db.Model(&User{}).Updates(aliceUser)

	var dbAliceUser2 User
	db.Model(&User{}).First(&dbAliceUser2, aliceUser.ID)
	log.Printf("Alice user in db 2: %+v", dbAliceUser2)

	// 观察日志输出，我们会发现：
	// 		2026/02/17 22:43:26 Alice user in db 1: {ID:1 Name:Alice Age:18 Enabled:true}
	// 		2026/02/17 22:43:26 Alice user in db 2: {ID:1 Name:Alice Age:18 Enabled:true}
	// 第一次更新并未成功，Enabled 字段仍然是 true。原因是 Gorm 的 Updates 方法默认会忽略零值（zero value）的字段，而 bool 类型的零值是 false。
	// 因此，在第一次更新时，Enabled 字段被视为零值，Gorm 就没有将其包含在 UPDATE 语句中，导致数据库中的 Enabled 字段保持不变。
	// 第二次更新时，Enabled 字段被设置为 true，这不是零值，因此 Gorm 包含了该字段在 UPDATE 语句中，成功更新了数据库中的 Enabled 字段。
	// 这是 Gorm 的设计问题，https://gorm.io/zh_CN/docs/update.html：
	// 使用 struct 更新时, GORM 将只更新非零值字段。 你可能想用 map 来更新属性，或者使用 Select 声明字段来更新。

	// 自动获取 aliceUser 实体的 id 放到 update 语句中
	// 使用 Model(&User) 需要指定 id，Where("id = ?", aliceUser.ID)
	db.Model(&aliceUser).Updates(map[string]interface{}{"enabled": false})

	var dbAliceUser3 User
	db.Model(&User{}).First(&dbAliceUser3, aliceUser.ID)
	log.Printf("Alice user in db 3: %+v", dbAliceUser3)

	// 更新成功
	// 2026/02/17 22:53:07 Alice user in db 3: {ID:2 Name:Alice Age:18 Enabled:false}
}
