package entities

import "time"

type Payment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	LoanID    uint      `json:"loan_id"`
	Amount    float64   `json:"amount"`
	DueDate   time.Time `json:"due_date"`
	PaidDate  *time.Time `json:"paid_date,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}