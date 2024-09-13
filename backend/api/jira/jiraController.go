package jira

import (
	"encoding/json"
	"net/http"
	"unified-atlassian-platform/api/jira/jiraService"
)

// HandleRequests handles Jira-related API requests.
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetJiraTasks(w, r)
	case http.MethodPost:
		handleCreateJiraTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetJiraTasks fetches Jira tasks.
func handleGetJiraTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := jiraService.GetTasks()
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// handleCreateJiraTask creates a new Jira task.
func handleCreateJiraTask(w http.ResponseWriter, r *http.Request) {
	var task jiraService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := jiraService.CreateTask(task); err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
