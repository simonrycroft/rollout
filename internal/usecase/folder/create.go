package folder

import (
	"fmt"
	folderdomain "rollout/internal/domain/folder"
	folderrepo "rollout/internal/repository/folder"
)

type Create struct {
	repo *folderrepo.Repository
}

func NewCreate(repo *folderrepo.Repository) *Create {
	return &Create{
		repo: repo,
	}
}

func (uc *Create) Execute(parentID uint, name string) (*folderdomain.Folder, error) {
	folder, err := folderdomain.NewFolder(parentID, name)
	if err != nil {
		return nil, fmt.Errorf("unable to create folder: %v", err)
	}

	if err := uc.repo.Save(folder); err != nil {
		return nil, fmt.Errorf("unable to create folder: %v", err)
	}

	return folder, nil
}
