package main

import (
	"net/http"

	"github.com/joaocvr/routes"
)

func main() {
	http.ListenAndServe(":8000", nil)
	routes.LoadRoutes()
}
