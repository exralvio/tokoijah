package handler

import "github.com/exralvio/tokoijah/models"

// JsonError is a generic error in JSON format
//
// swagger:response jsonError
type JsonMessage struct {
	// in: body
	Message string `json:"message"`
}

type ProductResponse struct {
	Data	[]models.Product		`json:"data"`
}

type OrderResponse struct {
	Data	[]models.Order		`json:"data"`
}

type PurchaseResponse struct {
	Data	[][]string		`json:"data"`
}
