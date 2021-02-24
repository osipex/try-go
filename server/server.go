package server

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e    *echo.Echo
	conn Storage
	cfg  *Config
}

func NewServer(cfg *Config) *Server {
	e := echo.New()

	return &Server{
		e:    e,
		conn: NewInMemoryStorage(),
		cfg:  cfg,
	}
}

func (s *Server) Run(host, port string) error {
	s.initRoutes()
	s.initMiddlewareLogging()
	return s.e.Start(host + ":" + port)
}

// Routes initialization
func (s *Server) initRoutes() {
	s.e.GET("/", s.HelloWorld)
	s.e.GET("/users", s.conn.GetAllUsers)
	s.e.POST("/users", s.conn.CreateUser)
	s.e.GET("/users/:id", s.conn.GetUser)
	s.e.DELETE("/users/:id", s.conn.DeleteUser)
	s.e.GET("/admin", s.Admin, middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(s.getUser())) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(s.getPass())) == 1 {
			return true, nil
		}
		return false, nil
	}))
}
