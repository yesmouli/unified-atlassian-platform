package jiraService

import "unified-atlassian-platform/db"

// Task represents a Jira task model.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// GetTasks retrieves all tasks from the database.
func GetTasks() ([]Task, error) {
	var tasks []Task
	query := "SELECT id, title, description, status FROM jira_tasks"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// CreateTask inserts a new task into the database.
func CreateTask(task Task) error {
	query := "INSERT INTO jira_tasks (title, description, status) VALUES (?, ?, ?)"
	_, err := db.DB.Exec(query, task.Title, task.Description, task.Status)
	return err
}
