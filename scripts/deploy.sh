#!/bin/bash

echo "Deploying application..."

# Pull latest code from repository
git pull origin main

# Build frontend
cd /frontend && npm install && npm run build

# Build Docker image for backend
cd /backend && docker build -t unified-atlassian-backend .

# Run Docker containers
docker-compose up -d

echo "Deployment complete."
