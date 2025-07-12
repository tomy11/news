package route

import (
	"news-api/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	auth := app.Group("/api/auth")
	
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}