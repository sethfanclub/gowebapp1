package routes

import (
	"fmt"
	"gowebapp1/app"
	"gowebapp1/models"
	"gowebapp1/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		username := r.PostForm.Get("username")
		var user models.User
		app.DB.First(&user, "username = ?", username)
		if user.ID > 0 {
			errorText := fmt.Sprintf("User '%s' already exists", username)
			http.Error(w, errorText, http.StatusBadRequest)
		}

		password1, password2 := r.PostForm.Get("password1"), r.PostForm.Get("password2")
		if password1 != password2 {
			http.Error(w, "Passwords didn't match", http.StatusBadRequest)
		}
		app.DB.Create(&models.User{Username: username, Password: password1})
		utils.LoginUser(w, r, user)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	utils.RenderTemplate(w, r, "register.html", map[string]interface{}{"Title": "Register"})
}
