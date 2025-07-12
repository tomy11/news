package app

import (
	"context"
	"errors"
	"time"
	"news-api/internal/domain/entity"
	"news-api/internal/domain/repository"
	"news-api/internal/domain/usecase"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type authUsecaseImpl struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(userRepo repository.UserRepository, jwtSecret string) usecase.AuthUsecase {
	return &authUsecaseImpl{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *authUsecaseImpl) Register(ctx context.Context, req usecase.RegisterRequest) (*usecase.AuthResponse, error) {
	existingUser, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	user := &entity.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := u.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &usecase.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (u *authUsecaseImpl) Login(ctx context.Context, req usecase.AuthRequest) (*usecase.AuthResponse, error) {
	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := u.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &usecase.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (u *authUsecaseImpl) generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtSecret))
}