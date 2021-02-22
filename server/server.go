package server

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e    *echo.Echo
	conn Storage
}

func NewServer(_ *Config) *Server {
	e := echo.New()

	return &Server{
		e:    e,
		conn: NewInMemoryStorage(),
	}

}

func (s *Server) Run() error {
	s.initRoutes()
	s.initMiddleware()
	return s.e.Start(":1323")
}

// Routes initialization
func (s *Server) initRoutes() {
	s.e.GET("/", s.HelloWorld)
	s.e.GET("/users", s.GetAllUsers)
	s.e.POST("/users", s.CreateUser)
	s.e.GET("/users/:id", s.GetUser)
	s.e.DELETE("/users/:id", s.DeleteUser)
	s.e.GET("/admin", s.Admin, middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("posypenko")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("123")) == 1 {
			return true, nil
		}
		return false, nil
	}))
}
