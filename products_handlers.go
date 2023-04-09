package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/uptrace/bunrouter"
)

func productsHandler(w http.ResponseWriter, req bunrouter.Request) error {
	db := ConnectDB()
	var products []Product
	db.Find(&products)
	return json.NewEncoder(w).Encode(products)
}

func newProductHandler(w http.ResponseWriter, req bunrouter.Request) error {
	http.ServeFile(w, req.Request, "form.html")
	return nil
}

func createProductHandler(w http.ResponseWriter, req bunrouter.Request) error {
	db := ConnectDB()
	req.ParseForm()
	title := req.Form.Get("title")
	price, _ := strconv.Atoi(req.Form.Get("price"))
	newProduct := Product{
		Title: title,
		Price: uint(price),
	}
	db.Create(&newProduct)
	// TODO validations
	var id string = strconv.FormatUint(uint64(newProduct.ID), 10)
	http.Redirect(w, req.Request, "/products/"+id, http.StatusSeeOther)
	return nil
}

func productHandler(w http.ResponseWriter, req bunrouter.Request) error {
	db := ConnectDB()
	id := req.Param("id")
	var product Product
	db.First(&product, id)
	return json.NewEncoder(w).Encode(product)
}
