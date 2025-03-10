package main

import (
	"fmt"
	"os"
)

func main() {
	// Create new storage instance
	storage, err := NewStorage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing storage: %v\n", err)
		os.Exit(1)
	}

	// Get command line arguments (excluding program name)
	args := os.Args[1:]

	// If no arguments provided, show usage
	if len(args) == 0 {
		showUsage()
		os.Exit(1)
	}

	// Execute the command
	if err := executeCommand(storage, args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// showUsage prints the help message
func showUsage() {
	fmt.Println(`Task Tracker CLI - Usage:

		Commands:
		add <description>                 Add a new task
		update <id> <description>         Update an existing task
		delete <id>                       Delete a task
		mark-in-progress <id>             Mark a task as in progress
		mark-done <id>                    Mark a task as done
		list                             List all tasks
		list todo                        List todo tasks
		list in-progress                 List in-progress tasks
		list done                        List completed tasks

		Examples:
		task-cli add "Buy groceries"
		task-cli update 1 "Buy groceries and cook dinner"
		task-cli mark-in-progress 1
		task-cli list
		`)
}
