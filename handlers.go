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

	c.JSON(http.StatusOK, products)
}

func readone(c *gin.Context) {

	db := connectDb()

	product_name := c.Param("product_name")

	var query string = fmt.Sprintf("SELECT * FROM products where product_name = '%s';", product_name)

	defer db.Close()
	row := db.QueryRow(query)

	var product Product
	err := row.Scan(&product.ProductName, &product.Price)

	checkErr(err)
	c.JSON(http.StatusOK, product)
}

func insertone(c *gin.Context) {

	db := connectDb()

	var product Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	defer db.Close()
	_, err := db.Exec("insert into products(product_name, price) values(?, ?);", product.ProductName, product.Price)
	checkErr(err)

	c.IndentedJSON(http.StatusCreated, product)
}

func updateone(c *gin.Context) {

	db := connectDb()

	var product Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	defer db.Close()
	_, err := db.Exec("update products set price = ? where product_name = ?", product.ProductName, product.Price)
	checkErr(err)

	c.JSON(http.StatusOK, product)
}

func deleteone(c *gin.Context) {

	db := connectDb()

	var product Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	defer db.Close()
	_, err := db.Exec("delete from products where product_name = ?", product.ProductName, product.Price)
	checkErr(err)

	c.JSON(http.StatusOK, product)
}
