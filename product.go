package models

import (
	"fmt"

	"github.com/joaocvr/golang-webstore/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Stock       int
}

func SearchAllProducts() []Product {
	db := db.ConnectDB()

	result, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)

	p := Product{}
	products := []Product{}

	for result.Next() {
		var id, stock int
		var name, description string
		var price float64

		err = result.Scan(&id, &name, &description, &price, &stock)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Id = id
		p.Description = description
		p.Price = price
		p.Stock = stock

		products = append(products, p)
	}

	defer db.Close()
	return products
}
