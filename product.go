package main

type Product struct {
	ID    uint   `gorm:"auto_random()" json:"id"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}
