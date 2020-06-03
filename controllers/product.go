package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/joaocvr/golang-webstore/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		stock := r.FormValue("stock")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Error at price convertion")
			fmt.Println(err)
		}

		stockInt, err := strconv.Atoi(stock)
		if err != nil {
			fmt.Println("Error at stock convertion")
			fmt.Println(err)
		}

		p := models.Product{}
		p.Name = name
		p.Description = description
		p.Price = priceFloat
		p.Stock = stockInt

		models.CreateNewProduct(p)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error at ID convertion")
		fmt.Println(err)
	}

	models.DeleteProduct(idInt)

	http.Redirect(w, r, "/", 301)
}
