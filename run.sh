#!/bin/bash

# run.sh - Wrapper script to run Go app with passed arguments

# Optional: check if at least one argument is passed
if [ $# -eq 0 ]; then
  echo "Usage: ./run.sh [args...]"
  echo "Example: ./run.sh workout 1"
  exit 1
fi

# Clear the terminal screen
clear

# Add a few empty lines after the clear command
echo -e "\n"

# Run the Go program with all provided arguments
go run main.go "$@"

# Add a few more empty lines after the output
# echo -e "\n"
