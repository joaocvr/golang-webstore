package models

import (
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

func CreateNewProduct(newProduct Product) {
	db := db.ConnectDB()

	insertNewProduct, err := db.Prepare("insert into products (name, description, price, stock) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertNewProduct.Exec(newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Stock)

	defer db.Close()
}

func DeleteProduct(id int) {
	db := db.ConnectDB()

	delete, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}
