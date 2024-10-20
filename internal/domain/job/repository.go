package job

// Repository defines the contract a repository must satisfy to store jobs
type Repository interface {
	Save(j *Job) error
	ListAll() ([]*Job, error)
}
