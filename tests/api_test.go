package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/w3aman/kong-service-catalog/api"
)

func TestGetServices(t *testing.T) {
	router := mux.NewRouter()
	api.RegisterRoutes(router)

	req, _ := http.NewRequest("GET", "/services?name=Service A&page=1&limit=1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Service A")
}

func TestGetServiceByID(t *testing.T) {
	router := mux.NewRouter()
	api.RegisterRoutes(router)

	req, _ := http.NewRequest("GET", "/services/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Service A")
}

func TestGetServiceVersions(t *testing.T) {
	router := mux.NewRouter()
	api.RegisterRoutes(router)

	req, _ := http.NewRequest("GET", "/services/1/versions", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "v1.0")
	assert.Contains(t, rr.Body.String(), "v2.0")
}
