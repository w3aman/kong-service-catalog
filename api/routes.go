package api

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes registers all the routes for the service catalog API
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/services", GetServices).Methods("GET")
	router.HandleFunc("/services/{id}", GetServiceByID).Methods("GET")
	router.HandleFunc("/services/{id}/versions", GetServiceVersions).Methods("GET")
}
