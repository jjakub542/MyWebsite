package database

import (
	"fmt"
	"log"

	"my-website-backend/internal/config"
	"my-website-backend/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to PostgreSQL database successfully")

	err = DB.AutoMigrate(&models.User{}, &models.Message{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	return DB
}

// MakeSuperuser creates a verified admin account if one doesn't exist
func MakeSuperuser(db *gorm.DB, email, rawPassword string) {
	var count int64
	db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		log.Printf("Superuser %s already exists.", email)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash superuser password: %v", err)
		return
	}

	admin := models.User{
		Email:    email,
		Password: string(hashedPassword),
		IsAdmin:  true,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("Failed to seed superuser: %v", err)
		return
	}

	log.Printf("Seeded initial superuser: %s", email)
}

// ClearData removes all rows from the tables and resets ID counters, but keeps the tables
func ClearData(db *gorm.DB) error {
	// CASCADE ensures any foreign-key dependencies are also cleared safely
	return db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
}
