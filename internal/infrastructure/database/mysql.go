package database

import (
	"news-api/internal/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDatabase(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}