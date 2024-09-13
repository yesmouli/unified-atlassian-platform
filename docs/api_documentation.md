# API Documentation

## Overview
This document provides information about the API endpoints available in the Unified Atlassian Platform backend.

## Authentication
- **POST /api/auth/login**: Logs in a user and returns a JWT token.
- **GET /api/auth/user**: Fetches user information (protected route).

## Jira
- **GET /api/jira/tasks**: Retrieves tasks from Jira.

## Confluence
- **GET /api/confluence/docs**: Retrieves documents from Confluence.

## Bitbucket
- **GET /api/bitbucket/repos**: Retrieves repositories from Bitbucket.

## Workflow Automation
- **GET /api/automation/workflows**: Retrieves workflows from the automation system.

## Error Handling
- The API responds with appropriate HTTP status codes and error messages in case of issues.
