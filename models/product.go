package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Sku				string
	Name			string
	Qty				int
}

func ProductMigration(db *gorm.DB){
	db.AutoMigrate(&Product{})
}

func GetOneProduct(db *gorm.DB, product_id int) Product {
	product := Product{}
	db.First(&product, product_id)

	return product
}