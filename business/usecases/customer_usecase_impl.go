package usecases

import (
	"context"
	"errors"
	"time"
	"news-api/business/entities"
	"news-api/business/repositories"
	"news-api/presentation/dto"
	"news-api/infra/external"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type customerUsecaseImpl struct {
	customerRepo      repositories.CustomerRepository
	creditScoreClient *external.CreditScoreClient
	jwtSecret         string
}

func NewCustomerUsecase(customerRepo repositories.CustomerRepository, jwtSecret string) CustomerUsecase {
	creditClient := external.NewCreditScoreClient("http://localhost:9001", "dev-api-key")
	return &customerUsecaseImpl{
		customerRepo:      customerRepo,
		creditScoreClient: creditClient,
		jwtSecret:         jwtSecret,
	}
}

func (u *customerUsecaseImpl) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existingCustomer, err := u.customerRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingCustomer != nil {
		return nil, errors.New("customer already exists")
	}

	customer := &entities.Customer{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Phone:    req.Phone,
	}

	if err := customer.HashPassword(); err != nil {
		return nil, err
	}

	if err := u.customerRepo.Create(ctx, customer); err != nil {
		return nil, err
	}

	token, err := u.generateToken(customer.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token:    token,
		Customer: u.mapToCustomerResponse(customer),
	}, nil
}

func (u *customerUsecaseImpl) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	customer, err := u.customerRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !customer.CheckPassword(req.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := u.generateToken(customer.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token:    token,
		Customer: u.mapToCustomerResponse(customer),
	}, nil
}

func (u *customerUsecaseImpl) GetProfile(ctx context.Context, customerID uint) (*dto.CustomerResponse, error) {
	customer, err := u.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	
	response := u.mapToCustomerResponse(customer)
	return &response, nil
}

func (u *customerUsecaseImpl) UpdateProfile(ctx context.Context, customerID uint, req dto.UpdateProfileRequest) (*dto.CustomerResponse, error) {
	customer, err := u.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	customer.Name = req.Name
	customer.Phone = req.Phone
	customer.Address = req.Address

	if err := u.customerRepo.Update(ctx, customer); err != nil {
		return nil, err
	}

	response := u.mapToCustomerResponse(customer)
	return &response, nil
}

func (u *customerUsecaseImpl) ChangePassword(ctx context.Context, customerID uint, req dto.ChangePasswordRequest) error {
	customer, err := u.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return err
	}

	if !customer.CheckPassword(req.CurrentPassword) {
		return errors.New("current password is incorrect")
	}

	customer.Password = req.NewPassword
	if err := customer.HashPassword(); err != nil {
		return err
	}

	return u.customerRepo.Update(ctx, customer)
}

func (u *customerUsecaseImpl) VerifyPhone(ctx context.Context, customerID uint, req dto.VerifyPhoneRequest) error {
	customer, err := u.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return err
	}

	// TODO: Implement OTP verification logic
	// For now, just mark as verified
	customer.Phone = req.Phone
	customer.IsPhoneVerified = true

	return u.customerRepo.Update(ctx, customer)
}

func (u *customerUsecaseImpl) VerifyIdentity(ctx context.Context, customerID uint, req dto.VerifyIdentityRequest) error {
	customer, err := u.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return err
	}

	// TODO: Implement identity verification logic
	customer.IDCard = &req.IDCard
	customer.IsVerified = true

	return u.customerRepo.Update(ctx, customer)
}

func (u *customerUsecaseImpl) GetCreditScore(ctx context.Context, customerID uint) (*dto.CreditScoreResponse, error) {
	customer, err := u.customerRepo.GetByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	score, err := u.creditScoreClient.GetCreditScore(customer)
	if err != nil {
		return nil, err
	}

	// Update credit score in database
	customer.CreditScore = score
	u.customerRepo.Update(ctx, customer)

	grade := u.calculateCreditGrade(score)

	return &dto.CreditScoreResponse{
		CustomerID:   customerID,
		CreditScore:  score,
		ScoreGrade:   grade,
		LastUpdated:  time.Now(),
	}, nil
}

func (u *customerUsecaseImpl) mapToCustomerResponse(customer *entities.Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          customer.ID,
		Email:       customer.Email,
		Name:        customer.Name,
		Phone:       customer.Phone,
		Address:     customer.Address,
		CreditScore: customer.CreditScore,
		IsVerified:  customer.IsVerified,
		CreatedAt:   customer.CreatedAt,
		UpdatedAt:   customer.UpdatedAt,
	}
}

func (u *customerUsecaseImpl) calculateCreditGrade(score int) string {
	switch {
	case score >= 800:
		return "A"
	case score >= 700:
		return "B"
	case score >= 600:
		return "C"
	default:
		return "D"
	}
}

func (u *customerUsecaseImpl) generateToken(customerID uint) (string, error) {
	claims := jwt.MapClaims{
		"customer_id": customerID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.jwtSecret))
}