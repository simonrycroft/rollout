package jobdomain

// Repository defines the contract for a storage layer
type Repository interface {
	Save(j *Job) error
	ListAll() ([]*Job, error)
}
