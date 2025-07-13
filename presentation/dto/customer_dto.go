package dto

import "time"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required,min=2"`
	Phone    string `json:"phone" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateProfileRequest struct {
	Name    string `json:"name" validate:"required,min=2"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}

type VerifyPhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
	OTP   string `json:"otp" validate:"required,len=6"`
}

type VerifyIdentityRequest struct {
	IDCard     string `json:"id_card" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	BirthDate  string `json:"birth_date" validate:"required"`
}

type AuthResponse struct {
	Token    string           `json:"token"`
	Customer CustomerResponse `json:"customer"`
}

type CustomerResponse struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	CreditScore int       `json:"credit_score"`
	IsVerified  bool      `json:"is_verified"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreditScoreResponse struct {
	CustomerID   uint   `json:"customer_id"`
	CreditScore  int    `json:"credit_score"`
	ScoreGrade   string `json:"score_grade"` // A, B, C, D
	LastUpdated  time.Time `json:"last_updated"`
}

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}