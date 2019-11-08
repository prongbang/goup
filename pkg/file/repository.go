package file

import (
	"mime/multipart"
)

const (
	imagePath = "public/assets/images"
)

// Repository is the interface
type Repository interface {
	SaveFile(m *multipart.FileHeader) (string, error)
}

type repository struct {
	FileDs DataSource
}

// NewRepository is the create repository
func NewRepository(datasource DataSource) Repository {
	return &repository{
		FileDs: datasource,
	}
}

func (r *repository) SaveFile(m *multipart.FileHeader) (string, error) {
	return r.FileDs.SaveFile(imagePath, m)
}
