package main

import (
	"fmt"
	"log"

	"rbac-go/config"
	"rbac-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err:= godotenv.Load(); err != nil{
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	app := fiber.New()
	app.Use(cors.New())

	routes.Routes(app)

	fmt.Println("listening on port 8080")
	log.Fatal(app.Listen(":8080"))
}