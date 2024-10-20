package job

import (
	"encoding/json"
	"net/http"
	"rollout/internal/usecase/job"
)

type Controller struct {
	jobCreator *job.Create
}

func NewController(jobCreator *job.Create) *Controller {
	return &Controller{jobCreator: jobCreator}
}

func (c *Controller) CreateJob(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newJob, err := c.jobCreator.Execute(request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newJob); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
