package database

import (
	"log"
	"test/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartPostgres() *gorm.DB {
	dbUrl := "postgres://pg:pass@localhost:5432/test"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})
	return db
}
