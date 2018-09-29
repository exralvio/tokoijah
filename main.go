package main

import (
	connect "github.com/exralvio/tokoijah/db"
	"github.com/exralvio/tokoijah/router"
	"github.com/rs/cors"
	"log"
	"net/http"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	connect.Migrate()

	// create router and start listen on port 8000
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8001", setupGlobalMiddleware(router)))
}
