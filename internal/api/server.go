package api

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"log"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func StartServer(config configs.AppConfig ) {
	app := fiber.New()
	// app.Get("/health", HealthCheckup)
	
	db,err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error %v\n",err)
	}
	log.Println("database got connected")

	// run migration
	db.AutoMigrate(&domain.User{})

	resHandler := &rest.RestHandler{
		App : app,
		DB: db,
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