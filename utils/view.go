package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

// parse layout.html in initiation of app?

func prepareDataDefaults(r *http.Request, data map[string]interface{}) (map[string]interface{}, error) {
	output := data
	currentUser := GetCurrentUser(r)
	var defaults = map[string]interface{}{"CurrentUser": currentUser}
	for k, v := range defaults {
		output[k] = v
	}
	return output, nil
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data map[string]interface{}) error {
	URL := fmt.Sprintf("templates/%s", templateName)
	t, err := template.ParseFiles("templates/layout.html", URL)
	if err != nil {
		return err
	}
	preparedData, err := prepareDataDefaults(r, data)
	if err != nil {
		return err
	}
	if err := t.ExecuteTemplate(w, templateName, preparedData); err != nil {
		return err
	}
	return nil
}
