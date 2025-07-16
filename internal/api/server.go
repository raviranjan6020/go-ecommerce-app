package api

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)
func StartServer(config configs.AppConfig ) {
	app := fiber.New()
	// app.Get("/health", HealthCheckup)
	resHandler := &rest.RestHandler{
		App : app,
	}
	setupRoutes(resHandler)
	app.Listen(config.ServerPort)
}
// func HealthCheckup(ctx *fiber.Ctx) error {
// 	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "I am running!",
// 	})
// }

func setupRoutes (rh *rest.RestHandler){
	//  user handler
	handlers.SetupUserRoutes(rh)
	// transaction
	// catalog
}