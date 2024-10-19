package jobrepository

import (
	"gorm.io/gorm"
	"rollout/internal/jobdomain"
)

// JobRepository handles job data persistence
type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db}
}

// Job defines the jobdomain database model
type Job struct {
	gorm.Model
	Name   string
	Status jobdomain.Status `json:"status"`
}

// Save saves a job
func (r JobRepository) Save(job *jobdomain.Job) error {
	r.db.Create(job)
	return nil
}

// ListAll lists all jobs
func (r JobRepository) ListAll() ([]*jobdomain.Job, error) {
	return []*jobdomain.Job{&jobdomain.Job{}}, nil
}

// maps a GORM model to a domain Job struct
func toDomain(j Job) *jobdomain.Job {
	return &jobdomain.Job{
		ID:     j.ID,
		Name:   j.Name,
		Status: j.Status,
	}
}

// maps a domain Job struct to a GORM model
func toGorm(j *jobdomain.Job) *Job {
	return &Job{
		Model:  gorm.Model{ID: j.ID},
		Name:   j.Name,
		Status: j.Status,
	}
}
