package mysql

import (
	"context"
	"news-api/business/entities"
	"news-api/business/repositories"
	"gorm.io/gorm"
)

type mysqlCustomerRepository struct {
	db *gorm.DB
}

func NewMysqlCustomerRepository(db *gorm.DB) repositories.CustomerRepository {
	return &mysqlCustomerRepository{
		db: db,
	}
}

func (r *mysqlCustomerRepository) Create(ctx context.Context, customer *entities.Customer) error {
	return r.db.WithContext(ctx).Create(customer).Error
}

func (r *mysqlCustomerRepository) GetByEmail(ctx context.Context, email string) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *mysqlCustomerRepository) GetByID(ctx context.Context, id uint) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *mysqlCustomerRepository) Update(ctx context.Context, customer *entities.Customer) error {
	return r.db.WithContext(ctx).Save(customer).Error
}

func (r *mysqlCustomerRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entities.Customer{}, id).Error
}