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
		product_purchase := models.PurchaseByProductId(db, int(product.ID))
		average := models.SumBuyPrice(product_purchase)

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

	var sales []models.SaleItem
	var sale models.SaleItem
	for _, order := range orders {
		product_detail := models.GetOneProduct(db, order.ProductID)
		product_purchase := models.PurchaseByProductId(db, order.ProductID)
		buyprice := models.SumBuyPrice(product_purchase)

		sale.Invoice = order.Invoice
		sale.Date = fmt.Sprint(order.CreatedAt.Format("2006-01-02 15:04:05"))
		sale.Sku = product_detail.Sku
		sale.Name = product_detail.Name
		sale.Qty = order.Qty
		sale.SalePrice = int(order.Price)
		sale.Total = int(order.Total)
		sale.BuyPrice = int(math.Round(buyprice))
		profit := order.Total - (buyprice * float64(order.Qty))
		sale.Profit = int(math.Round(profit))
		sales = append(sales, sale)
	}

	json.NewEncoder(w).Encode(sales)
}