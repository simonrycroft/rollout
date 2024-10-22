package folder

import (
	"encoding/json"
	"net/http"
	"rollout/internal/usecase/folder"
)

type Controller struct {
	folderCreator *folder.Create
}

func NewController(folderCreator *folder.Create) *Controller {
	return &Controller{
		folderCreator: folderCreator,
	}
}

// TODO why is the response a value but the request a pointer?
func (c *Controller) CreateFolder(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ParentID uint   `json:"parent_id"`
		Name     string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
	}

	newFolder, err := c.folderCreator.Execute(request.ParentID, request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newFolder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
