package main

import (
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
	//Atualizar um produto
	//var product Product
	// db.First(&product, 1)
	// product.Name = "Data show"
	// db.Save(&product)

	// var product2 Product
	// db.First(&product2, 1)
	// fmt.Println(product2.Name)

	//Atualizando multiplas colunas
	db.Model(&Product{}).Where("id = ?", 1).
		Updates(Product{Price: 5000.00, Name: "Xbox Series X"})

	//Exclusão de produto
	var product3 Product
	db.First(&product3, 4)
	db.Delete(&product3)
}
