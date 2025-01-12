package config

import (
	"fmt"
	"log"
	"os"
	"rbac-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB*gorm.DB

func ConnectDB(){
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to database")

	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
	fmt.Println("Database Migrated")
}