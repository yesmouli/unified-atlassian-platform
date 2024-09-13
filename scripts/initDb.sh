#!/bin/bash

echo "Initializing database..."

# Example for initializing a PostgreSQL database
psql -U $DB_USER -d $DB_NAME -a -f /backend/db/migrations.sql

echo "Database initialization complete."
