package usecases

import "news-api/business/entities"

type PaymentUsecase interface {
	CreatePayment(payment *entities.Payment) error
	GetPaymentByID(id uint) (*entities.Payment, error)
	GetPaymentsByLoan(loanID uint) ([]*entities.Payment, error)
	ProcessPayment(paymentID uint) error
}