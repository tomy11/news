package mysql

import (
	"news-api/business/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDatabase(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.Customer{}, &entities.Loan{}, &entities.Payment{})
	if err != nil {
		return nil, err
	}

	return db, nil
}