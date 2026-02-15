package gorm

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dbPath string, entities []any, config *gorm.Config) *gorm.DB {

	db, err := gorm.Open(sqlite.Open(dbPath), config)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	for _, entity := range entities {
		if err := db.AutoMigrate(entity); err != nil {
			log.Fatalf("failed to auto migrate: %v", err)
		}
	}

	return db
}
