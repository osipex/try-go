package server

import (
	"github.com/labstack/echo/v4/middleware"
)

// TODO: IMPLEMENT PART B

func (s *Server) initMiddleware() {
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${level} | method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	// s.e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	// Be careful to use constant time comparison to prevent timing attacks
	// 	if subtle.ConstantTimeCompare([]byte(username), []byte("posypenko")) == 1 &&
	// 		subtle.ConstantTimeCompare([]byte(password), []byte("123")) == 1 {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))
}
