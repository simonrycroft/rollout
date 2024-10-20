package job

import (
	"gorm.io/gorm"
	jobdomain "rollout/internal/domain/job"
)

// Repository handles job storage
type Repository struct {
	db *gorm.DB
}

// NewRepository creates a new job repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

// Job defines the job database model
type Job struct {
	gorm.Model
	Name   string
	Status jobdomain.Status `json:"status"`
}

// Save saves a job
func (r Repository) Save(job *jobdomain.Job) error {
	// TODO creating vs updating
	r.db.Create(toGorm(job))
	return nil
}

// ListAll lists all jobs
func (r Repository) ListAll() ([]*jobdomain.Job, error) {
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
