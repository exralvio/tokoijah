package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	ProductID		uint		`json:"product_id", omitempty`
	Invoice			string
	Product			Product		`gorm:"foreignkey:ProductID"`
	Purchases		[]*Purchase	`gorm:"foreignkey:ProductID;association_foreignkey:ProductID"`
	Qty				int
	Price			float64
	Total			float64
	Note			string
}

func OrderMigration(db *gorm.DB){
	db.AutoMigrate(&Order{})
}