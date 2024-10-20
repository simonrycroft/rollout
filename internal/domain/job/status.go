package job

const (
	StatusPending Status = iota
	StatusRunning
	StatusSucceeded
	StatusFailed
	StatusCancelled
)

// Status represents a job status
type Status int

// String returns a string representation of a job status
func (s Status) String() string {
	return [...]string{"Pending", "Running", "Succeeded", "Failed", "Cancelled"}[s]
}

// IsValid checks that a job status is valid
func (s Status) IsValid() bool {
	return s >= StatusPending && s <= StatusCancelled
}
