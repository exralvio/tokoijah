package handler

import (
	"encoding/json"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreatePurchase(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var purchase models.Purchase
	err = json.NewDecoder(r.Body).Decode(&purchase)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	product_id := purchase.ProductID
	number_order := purchase.NumberOrder
	number_receive := purchase.NumberReceive
	price := purchase.Price
	total := float64(number_receive) * price
	note := purchase.Note
	receipt := purchase.Receipt

	db.Create(&models.Purchase{
			ProductID: product_id,
			NumberOrder: number_order,
			NumberReceive: number_receive,
			Price: price,
			Total: total,
			Note: note,
			Receipt: receipt,
	})

	// Update product stock
	var product models.Product
	db.First(&product, product_id)
	product.Qty = product.Qty + number_receive
	db.Save(&product)

	json.NewEncoder(w).Encode(JsonMessage{"Success adding purchase!"})
}
