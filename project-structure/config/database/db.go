package database

import (
	"fmt"
	"log"
	"os"
	"project-structure/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database ulanmadi: %v", err)
	}

	log.Println("✅ Database ulandi!")

	// AutoMigrate jadvallarni avtomatik yaratadi
	db.AutoMigrate(&models.User{})

	return db
}
