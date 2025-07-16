package api

import (
	"go-ecommerce-app/configs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)
func StartServer(config configs.AppConfig ) {
	app := fiber.New()
	app.Get("/health", HealthCheckup)
	app.Listen(config.ServerPort)
}
func HealthCheckup(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I am running!",
	})
}