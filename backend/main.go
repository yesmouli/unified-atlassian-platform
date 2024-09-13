package main

import (
	"log"
	"net/http"
	"unified-atlassian-platform/api/jira"
	"unified-atlassian-platform/api/confluence"
	"unified-atlassian-platform/api/bitbucket"
	"unified-atlassian-platform/api/auth"
	"unified-atlassian-platform/api/automation"
	"unified-atlassian-platform/config"
	"unified-atlassian-platform/db"
	"unified-atlassian-platform/middleware"
)

func main() {
	// Load configuration settings
	config.LoadConfig()

	// Initialize and connect to the database
	db.ConnectDatabase()

	// Create a new HTTP router
	router := http.NewServeMux()

	// Register API routes
	router.HandleFunc("/api/jira", jira.HandleRequests)
	router.HandleFunc("/api/confluence", confluence.HandleRequests)
	router.HandleFunc("/api/bitbucket", bitbucket.HandleRequests)
	router.HandleFunc("/api/auth", auth.HandleRequests)
	router.HandleFunc("/api/automation", automation.HandleRequests)

	// Apply authentication middleware
	finalRouter := middleware.AuthMiddleware(router)

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", finalRouter))
}
