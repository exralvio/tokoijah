package handler

import (
	"encoding/json"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"math"
	"net/http"
	"strconv"
)

func AllPurchases(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var purchases []models.Purchase
	db.Find(&purchases)

	var datas = [][]string{}
	for _, purchase := range purchases {
		product_detail := models.GetOneProduct(db, purchase.ProductID)

		date := purchase.CreatedAt.Format("2006-01-02 15:04:05")
		sku := product_detail.Sku
		name := product_detail.Name
		number_order := strconv.Itoa(purchase.NumberOrder)
		number_receive := strconv.Itoa(purchase.NumberReceive)
		price := strconv.Itoa(int(math.Round(purchase.Price)))
		total := strconv.Itoa(int(math.Round(purchase.Total)))
		receipt := purchase.Receipt
		note := purchase.Note

		datas = append(datas, []string{
			date,
			sku,
			name,
			number_order,
			number_receive,
			price,
			total,
			receipt,
			note,
		})
	}

	json.NewEncoder(w).Encode(PurchaseResponse{Data: datas})
}

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
