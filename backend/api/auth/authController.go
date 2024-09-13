package auth

import (
	"encoding/json"
	"net/http"
	"unified-atlassian-platform/api/auth/authService"
)

// HandleRequests handles authentication-related API requests.
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleLogin(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleLogin processes user login requests.
func handleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials authService.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := authService.Login(credentials)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
