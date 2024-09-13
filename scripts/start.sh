#!/bin/bash

echo "Starting application..."

# Start the backend server
cd /backend && npm start &

# Start the frontend server
cd /frontend && npm start

echo "Application started."
