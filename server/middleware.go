package server

import (
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func (s *Server) initMiddlewareLogging() {

	logging := s.cfg.Server.EnableLogging

	if logging {
		log.Printf("Logging enabled")
		s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} | ${level} method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
		}))
	} else {
		log.Printf("Logging disabled")
	}
}

func (s *Server) initMiddlewareAuth() {
	return
}
