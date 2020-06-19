package pkg

import (
	"github.com/labstack/echo"
)

// Server ...
type Server struct {
}

func (s *Server) Mount(e *echo.Echo) {
	e.Static("/", "frontend/dist/")
}
