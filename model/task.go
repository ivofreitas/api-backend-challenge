package model

import "time"

type Task struct {
	ID          string    `json:"id"`
	Summary     string    `json:"summary" validate:"required"`
	PerformedBy string    `json:"-"`
	PerformedAt time.Time `json:"performed_at" validate:"required"`
}
