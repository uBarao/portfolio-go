package main

import (
	"net/http"
	"portfolio/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
