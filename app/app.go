package app

import (
	"encoding/gob"
	"gowebapp1/database"
	"gowebapp1/models"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	Secrets map[string]string
	DB      *gorm.DB
	Store   *sessions.CookieStore
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	var err error
	DB, err = database.Init()
	if err != nil {
		return err
	}

	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	gob.Register(models.User{})

	return nil
}
