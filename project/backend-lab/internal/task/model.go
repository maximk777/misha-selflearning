package task

import (
	"errors"
	"time"
)

var (
	ErrNotFound      = errors.New("task not found")
	ErrConflict      = errors.New("task state conflict")
	ErrInvalidTitle  = errors.New("title is required")
	ErrNotConfigured = errors.New("adapter is not configured")
)

type Status string

const (
	StatusOpen      Status = "open"
	StatusCompleted Status = "completed"
)

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Event struct {
	ID     string
	Type   string
	TaskID string
}
