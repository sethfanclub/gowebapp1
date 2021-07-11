package routes

import (
	"gowebapp1/utils"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	utils.LogoutUser(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
