package handler

import (
	"encoding/json"
	"fmt"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

func CreateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var order models.Order
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	var lastorder models.Order
	db.Last(&lastorder)
	last_id, err := strconv.Atoi(fmt.Sprint(lastorder.ID))
	if err != nil {

	}
	last_id = last_id + 1

	currentTime := time.Now()
	invoice := fmt.Sprintf("ID-%s-%d", currentTime.Format("20060102"), last_id)
	product_id := order.ProductID
	qty := order.Qty
	price := order.Price
	total := float64(qty) * price
	note := order.Note
	if note == "" {
		note = fmt.Sprintf("Pesanan %s", invoice)
	}

	db.Create(&models.Order{
		Invoice: invoice,
		ProductID: product_id,
		Qty: qty,
		Price: price,
		Total: total,
		Note: note,
	})

	// Update product stock
	var product models.Product
	db.First(&product, product_id)
	product.Qty = product.Qty - qty
	db.Save(&product)

	json.NewEncoder(w).Encode(JsonMessage{"Success adding order!"})
}
