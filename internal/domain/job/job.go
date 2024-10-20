package job

import (
	"errors"
)

// Job defines a job domain model
type Job struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

// NewJob creates a new job domain model
func NewJob(name string) (*Job, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	return &Job{
		Name:   name,
		Status: StatusPending,
	}, nil
}
