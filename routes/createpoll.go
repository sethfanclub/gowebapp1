package routes

import (
	"gowebapp1/app"
	"gowebapp1/models"
	"gowebapp1/utils"
	"net/http"
)

func CreatePoll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		question := r.PostForm.Get("question")
		answer1, answer2 := r.PostForm.Get("answer1"), r.PostForm.Get("answer2")
		var newQuestion = models.Question{Content: question}
		app.DB.Create(&newQuestion)
		app.DB.Create(&models.Answer{Content: answer1, QuestionID: newQuestion.ID})
		app.DB.Create(&models.Answer{Content: answer2, QuestionID: newQuestion.ID})
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	utils.RenderTemplate(w, r, "createpoll.html", map[string]interface{}{"Title": "Create Poll"})
}
