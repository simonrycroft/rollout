package job

import (
	"fmt"
	jobdomain "rollout/internal/domain/job"
	jobrepo "rollout/internal/repository/job"
)

type Create struct {
	repo *jobrepo.Repository
}

func NewCreate(repo *jobrepo.Repository) *Create {
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
