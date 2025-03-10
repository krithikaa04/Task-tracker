package main

import (
	"fmt"
	"strconv"
	"strings"
)

// executeCommand handles the execution of CLI commands
func executeCommand(storage *Storage, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no command provided")
	}

	command := args[0]
	switch command {
	case "add":
		if len(args) < 2 {
			return fmt.Errorf("description required for add command")
		}
		description := strings.Join(args[1:], " ")
		task, err := storage.AddTask(description)
		if err != nil {
			return err
		}
		fmt.Printf("Task added successfully (ID: %d)\n", task.ID)

	case "update":
		if len(args) < 3 {
			return fmt.Errorf("ID and description required for update command")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid task ID")
		}
		description := strings.Join(args[2:], " ")
		if err := storage.UpdateTask(id, description); err != nil {
			return err
		}
		fmt.Printf("Task %d updated successfully\n", id)

	case "delete":
		if len(args) < 2 {
			return fmt.Errorf("ID required for delete command")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid task ID")
		}
		if err := storage.DeleteTask(id); err != nil {
			return err
		}
		fmt.Printf("Task %d deleted successfully\n", id)

	case "mark-in-progress":
		if len(args) < 2 {
			return fmt.Errorf("ID required for mark-in-progress command")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid task ID")
		}
		if err := storage.MarkTaskInProgress(id); err != nil {
			return err
		}
		fmt.Printf("Task %d marked as in-progress\n", id)

	case "mark-done":
		if len(args) < 2 {
			return fmt.Errorf("ID required for mark-done command")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid task ID")
		}
		if err := storage.MarkTaskDone(id); err != nil {
			return err
		}
		fmt.Printf("Task %d marked as done\n", id)

	case "list":
		var tasks []Task
		if len(args) > 1 {
			switch args[1] {
			case "todo":
				tasks = storage.GetTasksByStatus(TodoStatus)
			case "in-progress":
				tasks = storage.GetTasksByStatus(InProgressStatus)
			case "done":
				tasks = storage.GetTasksByStatus(DoneStatus)
			default:
				return fmt.Errorf("invalid list filter: use 'todo', 'in-progress', or 'done'")
			}
		} else {
			tasks = storage.GetAllTasks()
		}
		printTasks(tasks)

	default:
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
}

// printTasks displays tasks in a formatted way
func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("\nTasks:")
	fmt.Println("----------------------------------------")
	for _, task := range tasks {
		fmt.Printf("ID: %d\n", task.ID)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Status: %s\n", task.Status)
		fmt.Printf("Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated: %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println("----------------------------------------")
	}
}
