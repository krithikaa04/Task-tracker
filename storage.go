package main

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
)

const fileName = "tasks.json"

// Storage handles all file operations for tasks
type Storage struct {
	tasks []Task
}

// NewStorage creates a new storage instance
func NewStorage() (*Storage, error) {
	s := &Storage{
		tasks: make([]Task, 0),
	}

	// Try to load existing tasks
	err := s.loadTasks()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return s, nil
}

// loadTasks reads tasks from the JSON file
func (s *Storage) loadTasks() error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.tasks) //converts json data to go slices
}

// saveTasks writes tasks to the JSON file
func (s *Storage) saveTasks() error {
	data, err := json.MarshalIndent(s.tasks, "", "    ") //converts task list to json format
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

// AddTask adds a new task and saves to storage
func (s *Storage) AddTask(description string) (Task, error) {
	// Generate new ID (max + 1)
	newID := 1
	if len(s.tasks) > 0 {
		sort.Slice(s.tasks, func(i, j int) bool {
			return s.tasks[i].ID < s.tasks[j].ID
		})
		newID = s.tasks[len(s.tasks)-1].ID + 1
	}

	task := NewTask(newID, description)
	s.tasks = append(s.tasks, task)

	if err := s.saveTasks(); err != nil {
		return Task{}, err
	}

	return task, nil
}

// UpdateTask updates an existing task
func (s *Storage) UpdateTask(id int, description string) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Update(description)
			return s.saveTasks()
		}
	}
	return errors.New("task not found")
}

// DeleteTask removes a task by ID
func (s *Storage) DeleteTask(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return s.saveTasks()
		}
	}
	return errors.New("task not found")
}

// GetAllTasks returns all tasks
func (s *Storage) GetAllTasks() []Task {
	return s.tasks
}

// GetTasksByStatus returns tasks filtered by status
func (s *Storage) GetTasksByStatus(status TaskStatus) []Task {
	var filtered []Task
	for _, task := range s.tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// MarkTaskInProgress marks a task as in-progress
func (s *Storage) MarkTaskInProgress(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].MarkInProgress()
			return s.saveTasks()
		}
	}
	return errors.New("task not found")
}

// MarkTaskDone marks a task as done
func (s *Storage) MarkTaskDone(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].MarkDone()
			return s.saveTasks()
		}
	}
	return errors.New("task not found")
}
