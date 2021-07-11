package routes

import (
	"gowebapp1/app"
	"gowebapp1/models"
	"gowebapp1/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Results(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionID, _ := strconv.Atoi(vars["ID"])
	var question models.Question
	app.DB.First(&question, questionID)
	var answers []models.Answer
	app.DB.Find(&answers, "question_id = ?", questionID)
	utils.RenderTemplate(w, r, "pollresults.html", map[string]interface{}{"Title": "Results", "Question": question, "Answers": answers})
}
