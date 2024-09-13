package automation

import (
	"encoding/json"
	"net/http"
	"unified-atlassian-platform/api/automation/workflowService"
)

// HandleRequests handles workflow automation-related API requests.
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleCreateWorkflow(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleCreateWorkflow creates a new workflow automation.
func handleCreateWorkflow(w http.ResponseWriter, r *http.Request) {
	var workflow workflowService.Workflow
	err := json.NewDecoder(r.Body).Decode(&workflow)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := workflowService.CreateWorkflow(workflow); err != nil {
		http.Error(w, "Error creating workflow", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
