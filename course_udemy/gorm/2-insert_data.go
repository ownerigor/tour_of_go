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
	//Migrar a tabela
	db.AutoMigrate(&Product{})

	//Inserir dados
	//db.Create(&Product{Name: "Playstation 5", Price: 3500})
	products := []Product{
		{Name: "Notebook", Price: 5000.00},
		{Name: "Tablet", Price: 1500.00},
		{Name: "Fogão", Price: 1000},
	}
	db.Create(&products)
}
