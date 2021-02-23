package server

type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		Host          string `yaml:"host"`
		User          string `yaml:"user"`
		Password      string `yaml:"password"`
		EnableLogging bool   `yaml:"logging"`
	} `yaml:"server"`
}

func (s *Server) getUser() string {
	return s.cfg.Server.User
}

func (s *Server) getPass() string {
	return s.cfg.Server.Password
}
