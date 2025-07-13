package entities

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Password       string    `json:"-" gorm:"not null"`
	Name           string    `json:"name" gorm:"not null"`
	Phone          string    `json:"phone" gorm:"unique"`
	Address        string    `json:"address"`
	CreditScore    int       `json:"credit_score" gorm:"default:0"`
	IsVerified     bool      `json:"is_verified" gorm:"default:false"`
	IsPhoneVerified bool     `json:"is_phone_verified" gorm:"default:false"`
	IDCard         *string   `json:"-" gorm:"unique"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (c *Customer) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)
	return nil
}

func (c *Customer) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
	return err == nil
}