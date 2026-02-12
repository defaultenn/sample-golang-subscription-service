package storage

import (
	"test_task/internal/config"

	"gorm.io/gorm"
)

type Storage struct {
	common *gorm.DB
}

type StorageInterface interface {
	GetDatabase() *gorm.DB
}

func NewStorage(
	dbConfig config.DatabaseConfigInterface,
) StorageInterface {

	db, err := NewDatabase(dbConfig)

	if err != nil {
		panic(err)
	}

	return &Storage{
		common: db,
	}
}

func (s *Storage) GetDatabase() *gorm.DB {
	return s.common
}
