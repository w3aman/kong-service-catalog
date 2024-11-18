package models

// Service represents a service in the catalog
type Service struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Versions    []string `json:"versions"`
}
