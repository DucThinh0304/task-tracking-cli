package model

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int, name string) Task {
	now := time.Now()
	return Task{
		ID:          id,
		Description: name,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("Task ID: %d, Name: %s, Status: %s", t.ID, t.Description, t.Status)
}
