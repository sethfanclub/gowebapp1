package database

import (
	"fmt"
	"gowebapp1/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	fmt.Println("Connecting to DB")
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to DB")

	db.AutoMigrate(&models.User{}, &models.Question{}, &models.Answer{})

	return db, nil
}
