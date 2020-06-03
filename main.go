package main

import (
	"fmt"
	"net/http"

	"github.com/joaocvr/golang-webstore/routes"
)

func main() {
	fmt.Println("Starting server...")
	routes.LoadRoutes()
	fmt.Println("Routes loaded with success.")
	http.ListenAndServe(":8000", nil)
}
