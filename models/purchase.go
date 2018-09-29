package models

import (
	"github.com/jinzhu/gorm"
)

type Purchase struct {
	gorm.Model
	ProductID		int			`json:"product_id", omitempty`
	Product			Product		`gorm:"foreignkey:ProductID"`
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