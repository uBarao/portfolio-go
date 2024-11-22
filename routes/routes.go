package routes

import (
	"net/http"
	"portfolio/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Home)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
}
