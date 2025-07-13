package repositories

import "news-api/business/entities"

type PaymentRepository interface {
	Create(payment *entities.Payment) error
	GetByID(id uint) (*entities.Payment, error)
	GetByLoanID(loanID uint) ([]*entities.Payment, error)
	Update(payment *entities.Payment) error
	Delete(id uint) error
}