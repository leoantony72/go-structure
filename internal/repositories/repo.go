package repositories

import (
	"test/internal/ports"

	"gorm.io/gorm"
)

type TestRepo struct {
	db *gorm.DB
}

func NewTestRepo(db *gorm.DB) ports.TestRepo {
	return &TestRepo{
		db: db,
	}
}
