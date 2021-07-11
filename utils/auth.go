package utils

import (
	"fmt"
	"gowebapp1/app"
	"gowebapp1/models"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request, user models.User) error {
	LogoutUser(w, r)
	session, _ := app.Store.Get(r, "user")
	session.Values["userID"] = user.ID
	if err := session.Save(r, w); err != nil {
		return err
	}
	return nil
}

func LogoutUser(w http.ResponseWriter, r *http.Request) error {
	session, _ := app.Store.Get(r, "user")
	delete(session.Values, "userID")
	if err := session.Save(r, w); err != nil {
		return err
	}
	return nil
}

func GetCurrentUser(r *http.Request) *models.User {
	session, _ := app.Store.Get(r, "user")
	if len(session.Values) != 1 {
		return &models.User{}
	}
	userID := session.Values["userID"].(uint)
	if userID == 0 {
		return &models.User{}
	}
	user, err := GetUserByID(userID)
	if err != nil {
		return &models.User{}
	}
	return user
}

func GetUserByID(ID uint) (user *models.User, err error) {
	app.DB.First(&user, ID)
	if user.ID == 0 {
		return nil, err
	}
	return user, nil
}

func PrintUserSession(r *http.Request) {
	session, _ := app.Store.Get(r, "user")
	fmt.Println(session.Values)
}
