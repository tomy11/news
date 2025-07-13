package router

import (
	"news-api/presentation/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func CustomerRoutes(app *fiber.App, customerHandler *handlers.CustomerHandler) {
	// Auth routes (public)
	auth := app.Group("/api/auth")
	auth.Post("/register", customerHandler.Register)
	auth.Post("/login", customerHandler.Login)
	
	// Customer profile routes (protected)
	customers := app.Group("/api/customers")
	// TODO: Add JWT middleware here
	customers.Get("/profile", customerHandler.GetProfile)
	customers.Put("/profile", customerHandler.UpdateProfile)
	customers.Put("/password", customerHandler.ChangePassword)
	
	// Verification routes (protected)
	customers.Post("/verify-phone", customerHandler.VerifyPhone)
	customers.Post("/verify-identity", customerHandler.VerifyIdentity)
	customers.Get("/credit-score", customerHandler.GetCreditScore)
}