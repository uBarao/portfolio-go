package controllers

import (
	"html/template"
	"net/http"
	"portfolio/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	projects := []models.Project{
		{
			Name:         "Portfolio",
			Summary:      "Esta página, para mostrar mais sobre mim.",
			Description:  "Usando os templates em Golang e HTML e CSS para a estrutura das páginas, eu criei este site e o coloquei online a disposição para quem queira ver.",
			Technologies: []string{"Golang", "HTML", "CSS"},
			Image:        "/static/images/project-portfolio.png",
			Url:          "https://github.com/uBarao/portfolio-go",
		},
		{
			Name:         "Sistema de Sorteio",
			Summary:      "Um sistema que cria, registra e realiza sorteios.",
			Description:  "Usando todas as possibilidades de Go, eu criei o projeto que cria, recebe, exibe e realiza sorteios, além de transformar em uma API Rest.",
			Technologies: []string{"Golang"},
			Image:        "",
			Url:          "https://github.com/uBarao/raffle_system",
		},
	}

	temp.ExecuteTemplate(w, "Home", projects)
}
