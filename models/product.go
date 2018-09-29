package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Sku				string
	Name			string
	Qty				int
	Purchases		[]*Purchase		`gorm:"foreignkey:ProductID"`
}

func ProductMigration(db *gorm.DB){
	db.AutoMigrate(&Product{})
}