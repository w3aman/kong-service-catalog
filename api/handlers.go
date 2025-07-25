package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/w3aman/kong-service-catalog/repository"
)

// GetServices handles fetching a list of services with pagination, filtering, and sorting
func GetServices(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort_by")

	services, err := repository.GetServices(name, page, limit, sortBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

// GetServiceByID handles fetching a specific service by its ID
func GetServiceByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	service, err := repository.GetServiceByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

// GetServiceVersions handles fetching all versions of a specific service
func GetServiceVersions(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	versions, err := repository.GetServiceVersions(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(versions)
}
