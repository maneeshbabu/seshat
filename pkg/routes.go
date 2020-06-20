package pkg

import (
	v1 "github.com/amagimedia/seshat/handlers/v1/api"
	"github.com/amagimedia/seshat/repository"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

// Server ...
type Server struct {
}

// Mount mounts all the routes
func (s *Server) Mount(e *echo.Echo) {
	e.Static("/", "frontend/dist/")
	s.registerV1Routes(e)
}

func (s *Server) registerV1Routes(e *echo.Echo) {
	e.Validator = &repository.CustomValidator{Validator: validator.New()}
	g := e.Group("/v1/api")
	g.GET("/agents", v1.ListAgent)
	g.POST("/agents", v1.CreateAgent)
}
