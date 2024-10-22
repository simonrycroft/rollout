package folder

import (
	"gorm.io/gorm"
	folderdomain "rollout/internal/domain/folder"
)

type Folder struct {
	gorm.Model
	ParentID uint
	Name     string
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r Repository) Save(folder *folderdomain.Folder) error {
	// TODO creating vs updating
	r.db.Create(toGorm(folder))
	return nil
}

func toGorm(f *folderdomain.Folder) *Folder {
	return &Folder{
		Model:    gorm.Model{ID: f.ID},
		ParentID: f.ParentID, //todo optional
		Name:     f.Name,
	}
}
