package routes

import (
	"gowebapp1/app"
	"gowebapp1/models"
	"gowebapp1/utils"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var questions []models.Question
	app.DB.Find(&questions)
	utils.RenderTemplate(w, r, "index.html", map[string]interface{}{"Title": "Polls", "Questions": questions})
}
