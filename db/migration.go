package config

import (
	"fmt"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Migrate(){
	db, err = gorm.Open("sqlite3", "tokoijah.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	models.ProductMigration(db)
	models.PurchaseMigration(db)
	models.OrderMigration(db)

	defer db.Close()
}