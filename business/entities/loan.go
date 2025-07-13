package entities

import "time"

type Loan struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CustomerID  uint      `json:"customer_id"`
	Amount      float64   `json:"amount"`
	InterestRate float64  `json:"interest_rate"`
	Term        int       `json:"term"` // in months
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}