package main

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api"
	"log"
)

func main() {
	// fmt.Println("I am main function")
	cfg, err := configs.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded %v\n", err)
	}
	api.StartServer(cfg)
}
