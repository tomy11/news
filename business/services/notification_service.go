package services

import "news-api/business/entities"

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) SendLoanApproval(customer *entities.Customer, loan *entities.Loan) error {
	// Send loan approval notification
	return nil
}

func (s *NotificationService) SendPaymentReminder(customer *entities.Customer, payment *entities.Payment) error {
	// Send payment reminder
	return nil
}