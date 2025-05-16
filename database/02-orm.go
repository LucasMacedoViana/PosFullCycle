package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product2 struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func ORM() {
	dsn := "root:root@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product2{})
	// Create
	//db.Create(&Product2{Name: "Dior", Price: 1000})

	// Create batch
	//db.Create([]Product2{
	//	{Name: "Chanel", Price: 2000},
	//	{Name: "Gucci", Price: 3000},
	//})

	// select one
	//var product Product2
	//db.First(&product, 1)
	//fmt.Println(product)
	//db.First(&product, "name = ?", "Chanel")
	//fmt.Println(product)

	// select all
	//var products []Product2
	//db.Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}
	// select limit 2
	//db.Limit(2).Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}
	// select offset 2 (paginação)
	//db.Limit(2).Offset(2).Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}

	//where
	//db.Where("name = ?", "Chanel").Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}
	//Like
	//db.Where("name LIKE ?", "%a%").Find(&products)
	//for _, p := range products {
	//	fmt.Println(p)
	//}

	//update
	//var p2 Product2
	//db.First(&p2, 1)
	//p2.Name = "New Dior"
	//db.Save(&p2)

	//delete
	//db.Delete(&p2)
}
