package repository

import (
	"errors"
	"strconv"

	"github.com/w3aman/kong-service-catalog/models"
)

// In-memory data storage for services
var services = []models.Service{
	{"1", "Service A", "Description for Service A", []string{"v1.0", "v2.0"}},
	{"2", "Service B", "Description for Service B", []string{"v1.1", "v1.2"}},
	{"3", "Service C", "Description for Service C", []string{"v1.0"}},
}

// GetServices retrieves services with pagination, filtering, and sorting
func GetServices(name, page, limit, sortBy string) ([]models.Service, error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Apply pagination
	start := (pageInt - 1) * limitInt
	end := start + limitInt
	if end > len(services) {
		end = len(services)
	}
	result := services[start:end]

	// Apply filtering by name
	if name != "" {
		filtered := []models.Service{}
		for _, service := range result {
			if service.Name == name {
				filtered = append(filtered, service)
			}
		}
		result = filtered
	}

	// Sorting is a placeholder, more logic can be added to sort by 'name' or other fields
	if sortBy == "name" {
		// Sorting logic (e.g., alphabetical sorting)
	}

	return result, nil
}

// GetServiceByID retrieves a specific service by its ID
func GetServiceByID(id string) (models.Service, error) {
	for _, service := range services {
		if service.ID == id {
			return service, nil
		}
	}
	return models.Service{}, errors.New("service not found")
}

// GetServiceVersions retrieves all versions of a specific service
func GetServiceVersions(id string) ([]string, error) {
	for _, service := range services {
		if service.ID == id {
			return service.Versions, nil
		}
	}
	return nil, errors.New("versions not found for service")
}
