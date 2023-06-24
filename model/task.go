package model

import "time"

type Task struct {
	ID          string    `json:"id"`
	Summary     string    `json:"summary"`
	PerformedAt time.Time `json:"performed_at"`
}
