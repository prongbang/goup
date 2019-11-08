package file

import (
	"log"
	"net/http"

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
	platform := c.FormValue("platform")
	log.Println(platform)

	filePath, err := h.UseCase.Upload(file)

	return c.JSON(http.StatusOK, echo.Map{
		"filePath": "http://localhost:4000/" + filePath,
	})
}
