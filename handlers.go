package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "active",
	})
}

func readAll(c *gin.Context) {

	db := connectDb()

	rows, err := db.Query("SELECT * FROM products;")
	checkErr(err)
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ProductName, &product.Price)
		checkErr(err)
		products = append(products, product)
	}

	fmt.Println(products)

	//product := Product{product_name: "burrito", price: 12.2}
	//fmt.Println(product)

	c.JSON(http.StatusOK, products)
}

func readone(c *gin.Context) {

	db := connectDb()

	product_name := c.Param("product_name")

	var query string = fmt.Sprintf("SELECT * FROM products where product_name = '%s';", product_name)
	rows, err := db.Query(query)
	checkErr(err)
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ProductName, &product.Price)
		checkErr(err)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func insertone(c *gin.Context) {

	db := connectDb()

	tx, err := db.Begin()
	checkErr(err)

	var product Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	stmt2, err := tx.Prepare("insert into products(product_name, price) values(?, ?);")
	checkErr(err)

	defer stmt2.Close()
	_, err = stmt2.Exec(product.ProductName, product.Price)
	checkErr(err)

	//c.JSON(http.StatusOK, product)
	c.IndentedJSON(http.StatusCreated, product)
}

func updateone(c *gin.Context) {

	db := connectDb()

	tx, err := db.Begin()
	checkErr(err)

	var product Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	stmt2, err := tx.Prepare("update products set price = ? where product_name = ?")
	checkErr(err)

	defer stmt2.Close()
	_, err = stmt2.Exec(product.ProductName, product.Price)
	checkErr(err)

	c.JSON(http.StatusOK, product)
}

func deleteone(c *gin.Context) {

	db := connectDb()

	tx, err := db.Begin()
	checkErr(err)

	var product Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	stmt2, err := tx.Prepare("delete from products where product_name = ?")
	checkErr(err)

	defer stmt2.Close()
	_, err = stmt2.Exec(product.ProductName, product.Price)
	checkErr(err)

	c.JSON(http.StatusOK, product)
}
