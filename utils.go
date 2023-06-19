package main

import "database/sql"

type Product struct {
	product_name string
	price        float32
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func connectDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./db1")
	checkErr(err)
	return db
}
