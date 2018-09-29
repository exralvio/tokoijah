package handler

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/exralvio/tokoijah/models"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	//"time"
)

func UploadFile(w http.ResponseWriter, r *http.Request){
	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%w", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	if mimeType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"{
		saveFile(w, file, handle)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JsonMessage{"Success!"})
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader){
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%w", err)
		return
	}

	err = ioutil.WriteFile("./import/"+handle.Filename, data, 0666)
	if err != nil {
		fmt.Fprintf(w, "%w", err)
		return
	}

	ImportXls(handle.Filename)
}


func ImportXls(filename string){
	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	xlsx, err := excelize.OpenFile("./import/"+filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows_product := xlsx.GetRows("Catatan Jumlah Barang")
	ImportProduct(rows_product)
	rows_purchase := xlsx.GetRows("Catatan Barang Masuk")
	ImportPurchase(rows_purchase)
	rows_orders := xlsx.GetRows("Catatan Barang Keluar")
	ImportOrder(rows_orders)

}

func ImportProduct(rows [][]string){
	n := 0
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if row[0] == "" || row[1] == "" {
			continue
		}

		qty_float,err := strconv.ParseFloat(row[2], 64)
		if err != nil {

		}
		qty := int(qty_float)

		db.Create(&models.Product{
			Sku: row[0],
			Name: row[1],
			Qty: qty,
		})

		n++
	}

	fmt.Println(strconv.Itoa(n) + " Product Imported")
}

func ImportPurchase(rows [][]string){
	n := 0

	product := models.Product{}
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if row[1] == "" || row[2] == "" {
			continue
		}

		sku := row[1]

		product = models.Product{}
		db.First(&product, "sku = ?", sku)

		if product.ID == 0 {
			continue
		}

		number_order_float, err := strconv.ParseFloat(row[3], 64)
		number_receive_float, err := strconv.ParseFloat(row[4], 64)
		price_float, err := strconv.ParseFloat(row[5], 64)
		total_float, err := strconv.ParseFloat(row[6], 64)
		if err != nil {

		}

		product_id := int(product.ID)
		number_order := int(number_order_float)
		number_receive := int(number_receive_float)
		price := float64(price_float)
		total := float64(total_float)
		receipt := row[7]
		note := row[8]

		db.Create(&models.Purchase{
			ProductID: product_id,
			NumberOrder: number_receive,
			NumberReceive: number_order,
			Price: price,
			Total: total,
			Receipt: receipt,
			Note: note,
		})

		// Update product stock
		product := models.Product{}
		db.First(&product, product_id)
		product.Qty = product.Qty + number_receive
		db.Save(&product)

		n++
	}

	fmt.Println(strconv.Itoa(n) + " Purchase Imported")
}

func ImportOrder(rows [][]string){
	n := 0
	max := len(rows)

	product := models.Product{}
	lastorder := models.Order{}
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if row[1] == "" || row[2] == "" {
			continue
		}

		sku := row[1]

		product = models.Product{}
		db.First(&product,"sku = ?", sku)

		if product.ID == 0 {
			continue
		}

		qty_float,err := strconv.ParseFloat(row[3], 64)
		price_float, err := strconv.ParseFloat(row[4], 64)
		total_float, err := strconv.ParseFloat(row[5], 64)
		if err != nil {

		}

		lastorder = models.Order{}
		db.Last(&lastorder)
		last_id, err := strconv.Atoi(fmt.Sprint(lastorder.ID))
		if err != nil {

		}
		last_id = last_id + 1

		currentTime := time.Now()
		invoice := fmt.Sprintf("ID-%s-%d", currentTime.Format("20060102"), last_id)
		product_id := int(product.ID)
		qty := int(qty_float)
		price := float64(price_float)
		total := float64(total_float)
		note := row[6]

		db.Create(&models.Order{
			//ProductID: product_id,
			Invoice: invoice,
			Qty: qty,
			Price: price,
			Total: total,
			Note: note,
		})

		n++

		// Update product stock
		product := models.Product{}
		db.First(&product, product_id)
		product.Qty = product.Qty - qty
		db.Save(&product)

		fmt.Println(strconv.Itoa(n) + " of " + strconv.Itoa(max) + " : " + product.Name + " / " + strconv.Itoa(product_id) + " / " + sku)
	}

	fmt.Println(strconv.Itoa(n) + " Order Imported")
}
