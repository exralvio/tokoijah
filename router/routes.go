package router

import (
	"github.com/exralvio/tokoijah/handlers"
	"net/http"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route {
		"GetAllProducts",
		"GET",
		"/products",
		handler.AllProducts,
	},
	Route {
		"GetOneProduct",
		"GET",
		"/products/{id}",
		handler.GetProduct,
	},
	Route {
		"CreateProduct",
		"POST",
		"/products",
		handler.CreateProduct,
	},
	Route {
		"UpdateProduct",
		"PUT",
		"/products/{id}",
		handler.UpdateProduct,
	},
	Route {
		"DeleteSku",
		"DELETE",
		"/products/{id}",
		handler.DeleteProduct,
	},
	{
		"CreatePurchase",
		"POST",
		"/purchase",
		handler.CreatePurchase,
	},
	{
		"CreateOrder",
		"POST",
		"/order",
		handler.CreateOrder,
	},
	{
		"ProductsReport",
		"GET",
		"/report/products",
		handler.ProductsReport,
	},
	{
		"SalesReport",
		"GET",
		"/report/sales",
		handler.SalesReport,
	},
	{
		"ExportProducts",
		"GET",
		"/export/products",
		handler.ExportProducts,
	},
	{
		"ExportPurchases",
		"GET",
		"/export/purchases",
		handler.ExportPurchases,
	},
	{
		"ExportOrders",
		"GET",
		"/export/orders",
		handler.ExportOrders,
	},
	{
		"ExportProductsReport",
		"GET",
		"/export/productsreport",
		handler.ExportProductsReport,
	},
	{
		"ExportSalesReport",
		"GET",
		"/export/salesreport",
		handler.ExportSalesReport,
	},
}
