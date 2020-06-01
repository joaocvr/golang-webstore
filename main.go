package main

import (
	"html/template"
	"net/http"

	"github.com/joaocvr/golang-webstore/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}
