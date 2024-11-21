package controllers

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Home", nil)
}
