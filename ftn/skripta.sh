#!/bin/sh
# Apply database migrations
echo "Create database migrations"
python3 manage.py makemigrations

echo "Apply database migrations"
python3 manage.py migrate

# Start server
echo "Starting server"
python3 manage.py runserver 0.0.0.0:8000