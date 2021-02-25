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

//
// Echo handler functions
//

func (s *Server) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is response from a web server!")
}

func (s *Server) GetAllUsers(c echo.Context) error {
	out, err := s.conn.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, out)
}

func (s *Server) CreateUser(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	err = s.conn.SaveUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusCreated, user)
}

func (s *Server) GetUser(c echo.Context) error {
	id := c.Param("id")
	out, err := s.conn.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, out)
}

func (s *Server) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := s.conn.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	return c.JSON(http.StatusOK, id)
}

func (s *Server) Admin(c echo.Context) error {
	return c.HTML(http.StatusOK, ""+
		"<html>\n\t<head>\n\t\t<style>body{background-color:#d0e4fe;}"+
		"h1{color:red;text-align:center;}p{font-family:\"Consolas\";font-size:22px;}</style>"+
		"\n\t</head>\n\t<body>\n\t\t<h1>WARNING ! ! !</h1>\n\t\t"+
		"<p>This is a TOP SECRET admin page</p>\n\t\t</body>\n\t</html>")

}
