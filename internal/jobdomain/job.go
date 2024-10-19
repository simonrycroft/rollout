package jobdomain

import "errors"

// Job defines a Rollout jobservice
type Job struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

func NewJob(name string) (*Job, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	return &Job{
		Name:   name,
		Status: StatusPending,
	}, nil
}
