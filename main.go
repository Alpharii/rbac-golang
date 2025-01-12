package main

import (
	"log"

	"rbac-go/config"
	"github.com/joho/godotenv"

)

func main() {
	if err:= godotenv.Load(); err != nil{
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
}