package usecases

import (
	"context"
	"errors"
	"time"
	"news-api/business/entities"
	"news-api/business/repositories"
	"github.com/golang-jwt/jwt/v5"
)

type loanUsecaseImpl struct {
	loanRepo     repositories.LoanRepository
	customerRepo repositories.CustomerRepository
	jwtSecret    string
}

type LoanUsecase interface {
	CreateLoan(ctx context.Context, loan *entities.Loan) error
	GetLoanByID(ctx context.Context, id uint) (*entities.Loan, error)
	GetCustomerLoans(ctx context.Context, customerID uint) ([]*entities.Loan, error)
	ApproveLoan(ctx context.Context, loanID uint) error
}

func NewLoanUsecase(loanRepo repositories.LoanRepository, customerRepo repositories.CustomerRepository, jwtSecret string) LoanUsecase {
	return &loanUsecaseImpl{
		loanRepo:     loanRepo,
		customerRepo: customerRepo,
		jwtSecret:    jwtSecret,
	}
}

func (u *loanUsecaseImpl) CreateLoan(ctx context.Context, loan *entities.Loan) error {
	customer, err := u.customerRepo.GetByID(ctx, loan.CustomerID)
	if err != nil {
		return err
	}
	if customer == nil {
		return errors.New("customer not found")
	}

	loan.Status = "pending"
	return u.loanRepo.Create(loan)
}

func (u *loanUsecaseImpl) GetLoanByID(ctx context.Context, id uint) (*entities.Loan, error) {
	return u.loanRepo.GetByID(id)
}

func (u *loanUsecaseImpl) GetCustomerLoans(ctx context.Context, customerID uint) ([]*entities.Loan, error) {
	return u.loanRepo.GetByCustomerID(customerID)
}

func (u *loanUsecaseImpl) ApproveLoan(ctx context.Context, loanID uint) error {
	loan, err := u.loanRepo.GetByID(loanID)
	if err != nil {
		return err
	}
	loan.Status = "approved"
	return u.loanRepo.Update(loan)
}

func (u *loanUsecaseImpl) generateToken(customerID uint) (string, error) {
	claims := jwt.MapClaims{
		"customer_id": customerID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtSecret))
}