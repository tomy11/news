package services

import "news-api/business/entities"

type LoanService struct{}

func NewLoanService() *LoanService {
	return &LoanService{}
}

func (s *LoanService) CalculateInterest(loan *entities.Loan) float64 {
	return loan.Amount * (loan.InterestRate / 100) / 12 * float64(loan.Term)
}

func (s *LoanService) ValidateLoanEligibility(customer *entities.Customer, amount float64) bool {
	// Business logic for loan eligibility
	return true
}