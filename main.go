package main

import (
	"net/http"

	"github.com/joaocvr/golang-webstore/routes"
)

func main() {
	http.ListenAndServe(":8000", nil)
	routes.LoadRoutes()
}
