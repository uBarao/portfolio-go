package main

import (
	"net/http"
	"portfolio/config"
	"portfolio/routes"
)

func main() {
	routes.LoadRoutes()
	port := config.GetPort()
	http.ListenAndServe(":"+port, nil)
}
