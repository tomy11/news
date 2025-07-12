package usecase

import (
	"context"
	"news-api/internal/domain/entity"
)

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  entity.User `json:"user"`
}

type AuthUsecase interface {
	Register(ctx context.Context, req RegisterRequest) (*AuthResponse, error)
	Login(ctx context.Context, req AuthRequest) (*AuthResponse, error)
}