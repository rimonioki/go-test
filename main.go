package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Loop through the data node for the FirstName

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		limit := r.URL.Query().Get("limit")
		products := getProducts(limit)
		json.NewEncoder(w).Encode(products)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getProducts(limit string) []Product {
	var result Response
	if limit == "" {
		limit = "5"
	}
	res, _ := Fetch("https://dummyjson.com/products?limit=" + string(limit))

	if err := json.Unmarshal(res, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	return result.Products
}
