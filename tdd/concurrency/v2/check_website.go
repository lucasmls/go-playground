package concurrency

import (
	"net/http"
)

// CheckWebsite returns true if a website returns 200 status code, false otherwise
func CheckWebsite(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}
