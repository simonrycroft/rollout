package folder

import "errors"

type Folder struct {
	ID       uint   `json:"id"`
	ParentID uint   `json:"parent_id"`
	Name     string `json:"name"`
}

func NewFolder(parentID uint, name string) (*Folder, error) {
	if name == "" {
		return nil, errors.New("folder name cannot be empty")
	}

	return &Folder{
		ParentID: parentID,
		Name:     name,
	}, nil
}
