package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running server at port :8080")

	r := gin.Default()

	r.GET("/healthcheck", healthCheck)

	r.GET("/product", readAll)

	r.GET("/product/:product_name", readone)

	r.POST("/product", insertone)

	r.PATCH("/product", updateone)

	r.DELETE("/product", deleteone)
	r.Run(":8080")
}
