package main

import (
	"time"
)

// TaskStatus represents the possible states of a task
type TaskStatus string

const (
	TodoStatus       TaskStatus = "todo"
	InProgressStatus TaskStatus = "in-progress"
	DoneStatus       TaskStatus = "done"
)

// Task represents a single task item
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// NewTask creates a new task with the given description
func NewTask(id int, description string) Task {
	now := time.Now()
	return Task{
		ID:          id,
		Description: description,
		Status:      TodoStatus,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Update modifies the task description and updates the UpdatedAt timestamp
func (t *Task) Update(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

// MarkInProgress changes the task status to in-progress
func (t *Task) MarkInProgress() {
	t.Status = InProgressStatus
	t.UpdatedAt = time.Now()
}

// MarkDone changes the task status to done
func (t *Task) MarkDone() {
	t.Status = DoneStatus
	t.UpdatedAt = time.Now()
}
