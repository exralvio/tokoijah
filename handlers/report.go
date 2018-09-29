package handler

import (
	"encoding/json"
	"fmt"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"math"
	"net/http"
)

func ProductsReport(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var products []models.Product
	db.Preload("Purchases").Find(&products)

	var product_items []models.ProductItem
	var product_item models.ProductItem
	for _, product := range products {
		average := models.SumBuyPrice(product.Purchases)

		product_item.Sku = product.Sku
		product_item.Name = product.Name
		product_item.Qty = product.Qty
		product_item.Average = int(math.Round(average))
		product_item.Total = int(math.Round((float64(product.Qty) * average)))

		product_items = append(product_items, product_item)
	}

	json.NewEncoder(w).Encode(product_items)
}

func SalesReport(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var orders []models.Order
	db.Find(&orders)

	json.NewEncoder(w).Encode(orders)
	return
	var sales []models.SaleItem
	var sale models.SaleItem
	for _, order := range orders {

		buyprice := models.SumBuyPrice(order.Purchases)

		sale.Invoice = order.Invoice
		sale.Date = fmt.Sprint(order.CreatedAt.Format("2006-01-02 15:04:05"))
		sale.Sku = order.Product.Sku
		sale.Name = order.Product.Name
		sale.Qty = order.Qty
		sale.SalePrice = order.Price
		sale.Total = order.Total
		sale.BuyPrice = buyprice
		sale.Profit = order.Total - (buyprice * float64(order.Qty))
		sales = append(sales, sale)
	}

	json.NewEncoder(w).Encode(sales)
}