package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter builds and returns a new router from routes
func NewRouter() *mux.Router {
	// When StrictSlash == true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	router := mux.NewRouter().StrictSlash(true)
	router.Use(Logger)

	/** Init Static **/
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)
	http.Handle("/", router)
	/** Init Static **/

	sub := router.PathPrefix("/v1").Subrouter()

	for _, route := range routes {
		sub.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)
	}

	return router
}
