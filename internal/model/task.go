package model

import "time"

type Task struct {
	ID        int
	Title     string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
