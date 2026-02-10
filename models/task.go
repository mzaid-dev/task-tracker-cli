package models

import "time"

// JSON keys are lowercase / snake_case
// Struct fields stay PascalCase

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskData struct {
	LastID int    `json:"last_id"`
	Tasks  []Task `json:"tasks"`
}
