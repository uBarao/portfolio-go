package routes

import (
	"net/http"
	"portfolio/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Home)
}
