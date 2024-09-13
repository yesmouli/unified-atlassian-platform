package utils

import (
	"net/http"
	"time"
)

// NewHTTPClient returns a new HTTP client with a custom timeout
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
