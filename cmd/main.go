package main

import (
	"log"
	"news-api/internal/app"
	"news-api/internal/delivery/http/handler"
	"news-api/internal/delivery/http/route"
	"news-api/internal/infrastructure/config"
	"news-api/internal/infrastructure/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewMysqlDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	userRepo := database.NewMysqlUserRepository(db)
	authUsecase := app.NewAuthUsecase(userRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authUsecase)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "News API is running",
			"version": "1.0.0",
		})
	})

	route.AuthRoutes(app, authHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(cfg.Port))
}