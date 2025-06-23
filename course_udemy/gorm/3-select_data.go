package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm: primaryKey`
	Name  string
	Price float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	//Selecionando 1 produto
	// var product Product
	// // db.First(&product, 2) //id
	// db.First(&product, "name = ?", "Fogão")
	// fmt.Println(&product)

	//Selecionando todos os produtos
	// var products []Product
	// //db.Find(&products)
	// db.Limit(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(&product)
	// }

	//Consulta com where
	var products []Product
	// db.Where("price >= ?", 3500).Find(&products)
	db.Where("name LIKE ?", "%Fo%").Find(&products)
	for _, product := range products {
		fmt.Println(&product)
	}
}
