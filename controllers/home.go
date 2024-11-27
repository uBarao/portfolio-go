package controllers

import (
	"html/template"
	"net/http"
	"portfolio/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	projects := []models.Project{
		{Name: "Portfolio", Summary: "", Description: "", Technologies: []string{"Golang", "HTML", "CSS"}, Image: "/static/images/project-portfolio.png"},
	}

	temp.ExecuteTemplate(w, "Home", projects)
}
