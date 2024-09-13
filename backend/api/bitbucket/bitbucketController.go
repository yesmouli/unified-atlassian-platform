package bitbucket

import (
	"encoding/json"
	"net/http"
	"unified-atlassian-platform/api/bitbucket/bitbucketService"
)

// HandleRequests handles Bitbucket-related API requests.
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRepositories(w, r)
	case http.MethodPost:
		handleCreateRepository(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetRepositories fetches Bitbucket repositories.
func handleGetRepositories(w http.ResponseWriter, r *http.Request) {
	repos, err := bitbucketService.GetRepositories()
	if err != nil {
		http.Error(w, "Error fetching repositories", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(repos)
}

// handleCreateRepository creates a new Bitbucket repository.
func handleCreateRepository(w http.ResponseWriter, r *http.Request) {
	var repo bitbucketService.Repository
	err := json.NewDecoder(r.Body).Decode(&repo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := bitbucketService.CreateRepository(repo); err != nil {
		http.Error(w, "Error creating repository", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
