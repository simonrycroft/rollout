package job

import (
	"fmt"
	jobdomain "rollout/internal/domain/job"
)

type Create struct {
	repo jobdomain.Repository
}

func NewCreate(repo jobdomain.Repository) *Create {
	return &Create{repo}
}

func (uc *Create) Execute(name string) (*jobdomain.Job, error) {
	job, err := jobdomain.NewJob(name)
	if err != nil {
		return nil, fmt.Errorf("unable to create job: %v", err)
	}

	if err = uc.repo.Save(job); err != nil {
		return nil, fmt.Errorf("unable to save job: %v", err)
	}

	return job, nil
}
