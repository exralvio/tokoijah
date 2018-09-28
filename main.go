package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/exralvio/tokoijah/db"
	model "github.com/exralvio/tokoijah/models"
	"github.com/exralvio/tokoijah/router"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	// populate our test database
	db.Insert(model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	db.Insert(model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})
	db.Insert(model.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	// create router and start listen on port 8000
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8001", setupGlobalMiddleware(router)))
}
