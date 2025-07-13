package mysql

import (
	"news-api/business/entities"
	"news-api/business/repositories"
	"gorm.io/gorm"
)

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) repositories.LoanRepository {
	return &loanRepository{db: db}
}

func (r *loanRepository) Create(loan *entities.Loan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepository) GetByID(id uint) (*entities.Loan, error) {
	var loan entities.Loan
	err := r.db.First(&loan, id).Error
	return &loan, err
}

func (r *loanRepository) GetByCustomerID(customerID uint) ([]*entities.Loan, error) {
	var loans []*entities.Loan
	err := r.db.Where("customer_id = ?", customerID).Find(&loans).Error
	return loans, err
}

func (r *loanRepository) Update(loan *entities.Loan) error {
	return r.db.Save(loan).Error
}

func (r *loanRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Loan{}, id).Error
}