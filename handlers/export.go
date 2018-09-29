package handler

import (
	"bytes"
	"encoding/csv"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

var datas = [][]string{}

func ExportProducts(w http.ResponseWriter, r *http.Request){
	date := time.Now()
	filename := "products-" + date.Format("20060102150405") + ".csv"

	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)

	/** Populating Data **/
	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var products []models.Product
	db.Find(&products)

	datas = datas[:0]
	for _, product := range products {
		datas = append(datas, []string{
			product.Sku,
			product.Name,
			strconv.Itoa(product.Qty),
		})
	}
	/** End populating Data **/

	for _, data := range datas {
		err := writer.Write(data)
		checkError("Cannot write to file", err)
	}

	writer.Flush()

	w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
	w.Header().Set("Content-Type", "text/csv")
	w.Write(b.Bytes())
}

func ExportPurchases(w http.ResponseWriter, r *http.Request){
	date := time.Now()
	filename := "purchases-" + date.Format("20060102150405") + ".csv"

	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)

	/** Populating Data **/
	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var purchases []models.Purchase
	db.Preload("Product").Find(&purchases)

	datas = datas[:0]
	for _, purchase := range purchases {
		datas = append(datas, []string{
			purchase.CreatedAt.Format("2006-01-02 15:04:05"),
			purchase.Product.Sku,
			purchase.Product.Name,
			strconv.Itoa(purchase.NumberOrder),
			strconv.Itoa(purchase.NumberReceive),
			strconv.Itoa(int(purchase.Price)),
			purchase.Receipt,
			purchase.Note,
		})
	}
	/** End populating Data **/

	for _, data := range datas {
		err := writer.Write(data)
		checkError("Cannot write to file", err)
	}

	writer.Flush()

	w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
	w.Header().Set("Content-Type", "text/csv")
	w.Write(b.Bytes())
}

func ExportOrders(w http.ResponseWriter, r *http.Request){
	date := time.Now()
	filename := "orders-" + date.Format("20060102150405") + ".csv"

	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)

	/** Populating Data **/
	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var orders []models.Order
	db.Preload("Product").Find(&orders)

	datas = datas[:0]
	for _, order := range orders {
		datas = append(datas, []string{
			order.CreatedAt.Format("2006-01-02 15:04:05"),
			order.Product.Sku,
			order.Product.Name,
			strconv.Itoa(order.Qty),
			strconv.Itoa(int(order.Price)),
			strconv.Itoa(int(order.Total)),
			order.Note,
		})
	}
	/** End populating Data **/

	for _, data := range datas {
		err := writer.Write(data)
		checkError("Cannot write to file", err)
	}

	writer.Flush()

	w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
	w.Header().Set("Content-Type", "text/csv")
	w.Write(b.Bytes())
}

func ExportProductsReport(w http.ResponseWriter, r *http.Request){
	date := time.Now()
	filename := "productsreport-" + date.Format("20060102150405") + ".csv"

	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)

	/** Populating Data **/
	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var products []models.Product
	db.Preload("Purchases").Find(&products)

	datas = datas[:0]
	for _, product := range products {
		average := models.SumBuyPrice(product.Purchases)

		datas = append(datas, []string{
			product.Sku,
			product.Name,
			strconv.Itoa(product.Qty),
			strconv.Itoa(int(average)),
			strconv.Itoa(int((float64(product.Qty) * average))),
		})
	}
	/** End populating Data **/

	for _, data := range datas {
		err := writer.Write(data)
		checkError("Cannot write to file", err)
	}

	writer.Flush()

	w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
	w.Header().Set("Content-Type", "text/csv")
	w.Write(b.Bytes())
}

func ExportSalesReport(w http.ResponseWriter, r *http.Request){
	date := time.Now()
	filename := "salesreport-" + date.Format("20060102150405") + ".csv"

	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)

	/** Populating Data **/
	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var orders []models.Order
	db.Preload("Product").Preload("Purchases").Find(&orders)

	datas = datas[:0]
	for _, order := range orders {
		buyprice := models.SumBuyPrice(order.Purchases)

		datas = append(datas, []string{
			order.Invoice,
			order.CreatedAt.Format("2006-01-02 15:04:05"),
			order.Product.Sku,
			order.Product.Name,
			strconv.Itoa(order.Qty),
			strconv.Itoa(int(order.Price)),
			strconv.Itoa(int(order.Total)),
			strconv.Itoa(int(buyprice)),
			strconv.Itoa(int(order.Total - (buyprice * float64(order.Qty)))),
		})
	}
	/** End populating Data **/

	for _, data := range datas {
		err := writer.Write(data)
		checkError("Cannot write to file", err)
	}

	writer.Flush()

	w.Header().Set("Content-Disposition", "attachment; filename=" + filename)
	w.Header().Set("Content-Type", "text/csv")
	w.Write(b.Bytes())
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}