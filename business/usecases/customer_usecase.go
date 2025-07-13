package usecases

import (
	"context"
	"news-api/presentation/dto"
)

type CustomerUsecase interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
	GetProfile(ctx context.Context, customerID uint) (*dto.CustomerResponse, error)
	UpdateProfile(ctx context.Context, customerID uint, req dto.UpdateProfileRequest) (*dto.CustomerResponse, error)
	ChangePassword(ctx context.Context, customerID uint, req dto.ChangePasswordRequest) error
	VerifyPhone(ctx context.Context, customerID uint, req dto.VerifyPhoneRequest) error
	VerifyIdentity(ctx context.Context, customerID uint, req dto.VerifyIdentityRequest) error
	GetCreditScore(ctx context.Context, customerID uint) (*dto.CreditScoreResponse, error)
}