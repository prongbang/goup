package file

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// Handler is the interface
type Handler interface {
	Upload(c echo.Context) error
}

type handler struct {
	UseCase UseCase
}

// NewHandler is new instance
func NewHandler(uc UseCase) Handler {
	return &handler{
		UseCase: uc,
	}
}

func (h *handler) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	filePath, err := h.UseCase.Upload(file)

	return c.JSON(http.StatusOK, echo.Map{
		"filePath": os.Getenv("HOST") + filePath,
	})
}
