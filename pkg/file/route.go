package file

import (
	"github.com/labstack/echo"
)

// Route is the interface
type Route interface {
	Initial(e *echo.Echo)
}

type route struct {
	Handle Handler
}

// NewRoute is new instance
func NewRoute(handle Handler) Route {
	return &route{
		Handle: handle,
	}
}

func (r *route) Initial(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	{
		v1.POST("/upload", r.Handle.Upload)
	}
}
