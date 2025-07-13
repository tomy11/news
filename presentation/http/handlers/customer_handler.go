package handlers

import (
	"strconv"
	"news-api/business/usecases"
	"news-api/presentation/dto"
	"news-api/presentation/validators"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerUsecase usecases.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecases.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: customerUsecase,
	}
}

// Auth endpoints
func (h *CustomerHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	response, err := h.customerUsecase.Register(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Registration failed",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ApiResponse{
		Success: true,
		Message: "Registration successful",
		Data:    response,
	})
}

func (h *CustomerHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	response, err := h.customerUsecase.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Success: false,
			Message: "Login failed",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiResponse{
		Success: true,
		Message: "Login successful",
		Data:    response,
	})
}

// Profile endpoints
func (h *CustomerHandler) GetProfile(c *fiber.Ctx) error {
	customerID := h.getCustomerIDFromToken(c)
	if customerID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Success: false,
			Message: "Unauthorized",
			Error:   "Invalid token",
		})
	}

	profile, err := h.customerUsecase.GetProfile(c.Context(), customerID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ApiResponse{
			Success: false,
			Message: "Profile not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiResponse{
		Success: true,
		Message: "Profile retrieved successfully",
		Data:    profile,
	})
}

func (h *CustomerHandler) UpdateProfile(c *fiber.Ctx) error {
	customerID := h.getCustomerIDFromToken(c)
	if customerID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Success: false,
			Message: "Unauthorized",
			Error:   "Invalid token",
		})
	}

	var req dto.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	profile, err := h.customerUsecase.UpdateProfile(c.Context(), customerID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Profile update failed",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiResponse{
		Success: true,
		Message: "Profile updated successfully",
		Data:    profile,
	})
}

func (h *CustomerHandler) ChangePassword(c *fiber.Ctx) error {
	customerID := h.getCustomerIDFromToken(c)
	if customerID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Success: false,
			Message: "Unauthorized",
			Error:   "Invalid token",
		})
	}

	var req dto.ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	err := h.customerUsecase.ChangePassword(c.Context(), customerID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Password change failed",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiResponse{
		Success: true,
		Message: "Password changed successfully",
	})
}

// Verification endpoints
func (h *CustomerHandler) VerifyPhone(c *fiber.Ctx) error {
	customerID := h.getCustomerIDFromToken(c)
	if customerID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Success: false,
			Message: "Unauthorized",
			Error:   "Invalid token",
		})
	}

	var req dto.VerifyPhoneRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	err := h.customerUsecase.VerifyPhone(c.Context(), customerID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Phone verification failed",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiResponse{
		Success: true,
		Message: "Phone verified successfully",
	})
}

func (h *CustomerHandler) VerifyIdentity(c *fiber.Ctx) error {
	customerID := h.getCustomerIDFromToken(c)
	if customerID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Success: false,
			Message: "Unauthorized",
			Error:   "Invalid token",
		})
	}

	var req dto.VerifyIdentityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := validators.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	err := h.customerUsecase.VerifyIdentity(c.Context(), customerID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Identity verification failed",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiResponse{
		Success: true,
		Message: "Identity verified successfully",
	})
}

func (h *CustomerHandler) GetCreditScore(c *fiber.Ctx) error {
	customerID := h.getCustomerIDFromToken(c)
	if customerID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Success: false,
			Message: "Unauthorized",
			Error:   "Invalid token",
		})
	}

	creditScore, err := h.customerUsecase.GetCreditScore(c.Context(), customerID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiResponse{
			Success: false,
			Message: "Failed to get credit score",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiResponse{
		Success: true,
		Message: "Credit score retrieved successfully",
		Data:    creditScore,
	})
}

// Helper method to extract customer ID from JWT token
func (h *CustomerHandler) getCustomerIDFromToken(c *fiber.Ctx) uint {
	// TODO: Implement JWT token extraction
	// For now, return a mock customer ID for testing
	customerIDStr := c.Get("X-Customer-ID", "1")
	customerID, err := strconv.ParseUint(customerIDStr, 10, 32)
	if err != nil {
		return 0
	}
	return uint(customerID)
}