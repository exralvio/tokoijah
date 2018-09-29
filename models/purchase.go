package models

import (
	"github.com/jinzhu/gorm"
)

type Purchase struct {
	gorm.Model
	ProductID		int			`json:"product_id", omitempty`
	NumberOrder		int			`json:"number_order", omitempty`
	NumberReceive	int			`json:"number_receive", omitempty`
	Price			float64
	Total			float64
	Receipt			string
	Note			string
}

func PurchaseMigration(db *gorm.DB){
	db.AutoMigrate(&Purchase{})
}

func PurchaseByProductId(db *gorm.DB, product_id int) []Purchase {
	purchases := []Purchase{}
	db.Find(&purchases, "product_id = ?", product_id)

	return purchases
}