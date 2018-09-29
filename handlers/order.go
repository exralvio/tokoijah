package handler

import (
	"encoding/json"
	"fmt"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"math"
	"net/http"
	"strconv"
	"time"
)

func AllOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var orders []models.Order
	db.Find(&orders)

	var datas = [][]string{}
	for _, order := range orders {
		product_detail := models.GetOneProduct(db, order.ProductID)

		date := order.CreatedAt.Format("2006-01-02 15:04:05")
		sku := product_detail.Sku
		name := product_detail.Name
		qty := strconv.Itoa(order.Qty)
		price := strconv.Itoa(int(math.Round(order.Price)))
		total := strconv.Itoa(int(math.Round(order.Total)))
		note := order.Note

		datas = append(datas, []string {
			date,
			sku,
			name,
			qty,
			price,
			total,
			note,
		})
	}

	json.NewEncoder(w).Encode(OrderResponse{Data: datas})
}

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
