package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Produto3 struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func ORM2() {
	dsn := "root:root@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Produto3{}, &Category{}, &SerialNumber{})

	//create category
	//category := Category{
	//	Name: "Eletrônicos",
	//}
	//db.Create(&category)

	//create product
	product := Produto3{
		Name:       "Notebook",
		Price:      2000,
		CategoryID: 1,
	}
	db.Create(&product)

	//create serial number
	serialNumber := SerialNumber{
		Number:    "123456",
		ProductID: 1,
	}
	db.Create(&serialNumber)

	var p3 []Produto3
	db.Preload("Category").Preload("SerialNumber").Find(&p3)
	for _, produto := range p3 {
		fmt.Println(produto.Name, produto.Category.Name, produto.SerialNumber.Number)
	}

}
