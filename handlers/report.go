package handler

import (
	"encoding/json"
	"fmt"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

type ProductItem struct {
	Sku			string		`json:"sku"`
	Name		string		`json:"name"`
	Qty			int			`json:"qty"`
	Average		float64		`json:"average"`
	Total		float64		`json:"total"`
}

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

	var product_items []ProductItem
	var product_item ProductItem
	for _, product := range products {
		average := SumBuyPrice(product.Purchases)

		product_item.Sku = product.Sku
		product_item.Name = product.Name
		product_item.Qty = product.Qty
		product_item.Average = average
		product_item.Total = (float64(product.Qty) * average)

		product_items = append(product_items, product_item)
	}

	json.NewEncoder(w).Encode(product_items)
}

type SaleItem struct {
	Invoice			string		`json:"invoice"`
	Date			string		`json:"date"`
	Sku				string		`json:"sku"`
	Name			string		`json:"name"`
	Qty				int			`json:"qty"`
	SalePrice		float64		`json:"sale_price"`
	Total			float64		`json:"total"`
	BuyPrice		float64		`json:"buy_price"`
	Profit			float64		`json:"profit"`
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
	db.Preload("Product").Preload("Purchases").Find(&orders)

	var sales []SaleItem
	var sale SaleItem
	for _, order := range orders {

		buyprice := SumBuyPrice(order.Purchases)

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

func SumBuyPrice(x[] *models.Purchase) float64{
	if len(x) == 0 {
		return 0
	}

	var total_price, result float64
	var total_receive int

	for _, value := range x {
		total_price += value.Total
		total_receive += value.NumberReceive
	}

	result = total_price / float64(total_receive)

	return result
}