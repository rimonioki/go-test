package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=myapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}

func CreateDummyData(db *gorm.DB) {
	var productsList []Product

	for i := 0; i < 10; i++ {
		productsList = append(productsList, generateProduct())
	}
	fmt.Println("Starting creating dummy products")
	db.CreateInBatches(productsList, 1000)
	fmt.Println("Finish creating dummy products")
}

func generateProduct() Product {
	title := faker.Name()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Product{
		Title: title,
		Price: uint(r1.Intn(100)),
	}
}
