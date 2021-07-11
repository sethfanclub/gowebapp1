package routes

import (
	"fmt"
	"gowebapp1/app"
	"gowebapp1/models"
	"gowebapp1/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Vote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionID, err := strconv.Atoi(vars["ID"])
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		answerID := r.PostForm.Get("answer")
		var answer models.Answer
		app.DB.First(&answer, answerID)
		answer.Votes += 1
		app.DB.Save(&answer)
		http.Redirect(w, r, fmt.Sprintf("/poll/%s/results", vars["ID"]), http.StatusFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var question models.Question
	app.DB.First(&question, questionID)
	var answers []models.Answer
	app.DB.Find(&answers, "question_id = ?", questionID)
	utils.RenderTemplate(w, r, "vote.html", map[string]interface{}{"Title": "Voting", "Question": question, "Answers": answers})
}
