package controllers

import (
	"net/http"
	"text/template"

	"github.com/joaocvr/golang-webstore/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func LoadRoutes() {
	http.HandleFunc("/", index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}
