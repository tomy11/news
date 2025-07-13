package repositories

import (
	"context"
	"news-api/business/entities"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *entities.Customer) error
	GetByEmail(ctx context.Context, email string) (*entities.Customer, error)
	GetByID(ctx context.Context, id uint) (*entities.Customer, error)
	Update(ctx context.Context, customer *entities.Customer) error
	Delete(ctx context.Context, id uint) error
}