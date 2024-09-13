package workflowService

import "unified-atlassian-platform/db"

// Workflow represents a workflow automation.
type Workflow struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateWorkflow inserts a new workflow into the database.
func CreateWorkflow(workflow Workflow) error {
	query := "INSERT INTO workflows (name, description) VALUES (?, ?)"
	_, err := db.DB.Exec(query, workflow.Name, workflow.Description)
	return err
}
