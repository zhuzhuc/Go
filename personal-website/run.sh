#!/bin/bash

# Check if PostgreSQL is running
pg_status=$(pg_ctl status 2>&1)
if [[ $pg_status == *"no server running"* ]]; then
  echo "Starting PostgreSQL..."
  pg_ctl start
fi

# Check if the database exists
if ! psql -lqt | cut -d \| -f 1 | grep -qw personal_website; then
  echo "Creating database personal_website..."
  createdb personal_website
fi

# Install dependencies
echo "Installing dependencies..."
go mod tidy

# Run the application
echo "Starting the application..."
go run main.go
