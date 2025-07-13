package repositories

import "news-api/business/entities"

type LoanRepository interface {
	Create(loan *entities.Loan) error
	GetByID(id uint) (*entities.Loan, error)
	GetByCustomerID(customerID uint) ([]*entities.Loan, error)
	Update(loan *entities.Loan) error
	Delete(id uint) error
}