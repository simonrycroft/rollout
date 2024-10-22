package folder

type FolderRepository interface {
	Save(f *Folder) error
	ListAll() ([]*Folder, error)
}
