package jobservice

import (
	"fmt"
	"rollout/internal/jobdomain"
)

type JobService struct {
	repo jobdomain.Repository
}

func NewJobService(repo jobdomain.Repository) *JobService {
	return &JobService{
		repo: repo,
	}
}

func (s *JobService) CreateJob(name string) (*jobdomain.Job, error) {
	newJob, err := jobdomain.NewJob(name)
	if err != nil {
		return nil, fmt.Errorf("unable to create job: %v", err)
	}

	err = s.repo.Save(newJob)
	if err != nil {
		return nil, fmt.Errorf("unable to create job: %v", err)
	}

	return newJob, nil
}
