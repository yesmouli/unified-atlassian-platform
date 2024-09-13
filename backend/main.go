package main

import (
    "log"
    "net/http"
    "unified-atlassian-platform/api/jira"
    "unified-atlassian-platform/api/confluence"
    "unified-atlassian-platform/api/bitbucket"
    "unified-atlassian-platform/api/auth"
    "unified-atlassian-platform/api/automation"
    "unified-atlassian-platform/db"
    "unified-atlassian-platform/middleware"
    "unified-atlassian-platform/config"
)

func main() {
    config.LoadConfig() // Load the configuration from config.yaml

    db.ConnectDatabase() // Connect to the database

    router := http.NewServeMux()

    // Add Routes
    router.HandleFunc("/api/jira", jira.HandleRequests)
    router.HandleFunc("/api/confluence", confluence.HandleRequests)
    router.HandleFunc("/api/bitbucket", bitbucket.HandleRequests)
    router.HandleFunc("/api/auth", auth.HandleRequests)
    router.HandleFunc("/api/automation", automation.HandleRequests)

    // Middleware
    wrappedRouter := middleware.AuthMiddleware(router)

    // Start the server
    log.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", wrappedRouter))
}
