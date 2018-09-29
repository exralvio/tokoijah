package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ProductID		int			`json:"product_id", omitempty`
	Invoice			string
	Qty				int
	Price			float64
	Total			float64
	Note			string
}

func OrderMigration(db *gorm.DB){
	db.AutoMigrate(&Order{})
}