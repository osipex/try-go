package server

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

// TODO: IMPLEMENT PART B

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func (s *Server) initMiddlewareLogging() {
	var c Config
	if err := cleanenv.ReadConfig("conf/conf.yaml", &c); err != nil {
		processError(err)
	}
	logging := c.Server.EnableLogging

	if logging {
		log.Printf("Logging enabled")
		s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339} | ${level} | method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
		}))
	} else {
		log.Printf("Logging disabled")
	}
}

func (s *Server) initMiddlewareAuth() {
	return
}
