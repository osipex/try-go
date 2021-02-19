package server

import (
	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func genUUID() string {
	id := guuid.New()
	return id.String()
}

func (s *Server) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is response from a web server!")
}
