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

	fileName, err := u.Repo.SaveFile(m)

	if err == nil {
		return fileName, err
	}

	return "", err
}
