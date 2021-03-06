package file

import (
	"mime/multipart"
)

// UseCase is the interface
type UseCase interface {
	Upload(m *multipart.FileHeader) (string, error)
}

type useCase struct {
	Repo Repository
}

// NewUseCase is the create usecase
func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}

func (u *useCase) Upload(m *multipart.FileHeader) (string, error) {
	return u.Repo.SaveFile(m)
}
