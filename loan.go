package main

import (
	"log"
	"news-api/infra/config"
	"news-api/infra/database/mysql"
	"news-api/presentation/http/handlers"
	"news-api/presentation/http/router"
	"news-api/business/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg := config.LoadConfig()

	db, err := mysql.NewMysqlDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	customerRepo := mysql.NewMysqlCustomerRepository(db)
	customerUsecase := usecases.NewCustomerUsecase(customerRepo, cfg.JWTSecret)
	customerHandler := handlers.NewCustomerHandler(customerUsecase)

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

	router.CustomerRoutes(app, customerHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(cfg.Port))
}
