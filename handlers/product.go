package handler

import (
	"encoding/json"
	"github.com/exralvio/tokoijah/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
)

var db *gorm.DB
var err error

func AllProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var products []models.Product
	db.Find(&products)

	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	db.Create(&models.Product{Sku: product.Sku, Name: product.Name, Qty: product.Qty})

	json.NewEncoder(w).Encode(JsonMessage{"Success Adding Product!"})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	vars := mux.Vars(r)
	product_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	var product models.Product
	db.First(&product, product_id)

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	db.Save(&product)

	json.NewEncoder(w).Encode(JsonMessage{"Product updated!"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	db, err = gorm.Open("sqlite3", "./tokoijah.db")
	if err != nil{
		panic("Could not connect to the datbase")
	}
	defer db.Close()

	vars := mux.Vars(r)
	product_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	var product models.Product
	db.Where("id = ?", product_id).Delete(&product)

	json.NewEncoder(w).Encode(JsonMessage{"Product Deleted!"})
}