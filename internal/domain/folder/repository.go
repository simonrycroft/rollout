package folder

type Repository interface {
	Save(f *Folder) error
	//ListAll() ([]*Folder, error)
}
