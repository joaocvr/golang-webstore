package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connection := "user=webstore dbname=webstore password=webstore00 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("ERROR connecting with data base.")
		fmt.Println(err)
		panic(err.Error())
	}

	return db
}
