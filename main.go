package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func main() {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
	)

	router.GET("/", indexHandler)
	db := ConnectDB()
	db.AutoMigrate(&Product{})
	CreateDummyData(db)
	router.GET("/products/new", newProductHandler)
	router.GET("/products", productsHandler)
	router.GET("/products/:id", productHandler)
	router.POST("/products", createProductHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
