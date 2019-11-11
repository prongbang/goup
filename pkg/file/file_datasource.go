package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// DataSource is the interface
type DataSource interface {
	SaveFile(path string, file *multipart.FileHeader) (string, error)
	CreateDirectory(dirName string) bool
	ExtWhitelist(path string) (string, error)
}

type dataSource struct {
}

// NewDataSource is new instance
func NewDataSource() DataSource {
	return &dataSource{}
}

func (ds *dataSource) SaveFile(path string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	if _, e := ds.ExtWhitelist(ext); e != nil {
		return "", e
	}

	uid, _ := uuid.NewUUID()
	filename := fmt.Sprintf("%s%s", uid.String(), ext)

	// Destination
	ds.CreateDirectory(path)
	pathFile := fmt.Sprintf("%s/%s", path, filename)
	dst, err := os.Create(pathFile)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return pathFile, err
}

func (ds *dataSource) CreateDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(err)
		}
		return true
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}

func (ds *dataSource) ExtWhitelist(path string) (string, error) {
	extList := os.Getenv("EXT")
	exts := strings.Split(extList, ",")

	if len(exts) > 0 {
		for _, ext := range exts {
			if filepath.Ext(path) == "."+ext {
				return "." + ext, nil
			}
		}
	}
	return "", fmt.Errorf("%s", "Unsupported file")
}
