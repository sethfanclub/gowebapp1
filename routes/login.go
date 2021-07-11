package routes

import (
	"fmt"
	"gowebapp1/app"
	"gowebapp1/models"
	"gowebapp1/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		username, password := r.PostForm.Get("username"), r.PostForm.Get("password")
		var user models.User
		app.DB.First(&user, "username = ?", username)
		if user.ID == 0 {
			errorString := fmt.Sprintf("User '%s' doesn't exist", username)
			http.Error(w, errorString, http.StatusBadRequest)
		}
		if password != user.Password {
			http.Error(w, "Incorrect password", http.StatusBadRequest)
		}
		if err := utils.LoginUser(w, r, user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	utils.RenderTemplate(w, r, "login.html", map[string]interface{}{"Title": "Login"})
}
