package confluence

import (
	"encoding/json"
	"net/http"
	"unified-atlassian-platform/api/confluence/confluenceService"
)

// HandleRequests handles Confluence-related API requests.
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetConfluencePages(w, r)
	case http.MethodPost:
		handleCreateConfluencePage(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetConfluencePages fetches Confluence pages.
func handleGetConfluencePages(w http.ResponseWriter, r *http.Request) {
	pages, err := confluenceService.GetPages()
	if err != nil {
		http.Error(w, "Error fetching pages", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pages)
}

// handleCreateConfluencePage creates a new Confluence page.
func handleCreateConfluencePage(w http.ResponseWriter, r *http.Request) {
	var page confluenceService.Page
	err := json.NewDecoder(r.Body).Decode(&page)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := confluenceService.CreatePage(page); err != nil {
		http.Error(w, "Error creating page", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
