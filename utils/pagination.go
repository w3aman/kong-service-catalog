package utils

import (
	"net/http"
	"strconv"
)

// ParsePaginationParams parses the pagination parameters from the query string
func ParsePaginationParams(r *http.Request) (page, limit int) {
	page, _ = strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	return page, limit
}
