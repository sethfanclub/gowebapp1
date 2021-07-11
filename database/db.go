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
	dbURL := fmt.Sprintf(
		"host=%s user=%s password=%s database=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to DB")

	db.AutoMigrate(&models.User{}, &models.Question{}, &models.Answer{})

	return db, nil
}
