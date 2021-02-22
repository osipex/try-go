package server

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		Host          string `yaml:"host"`
		User          string `yaml:"user"`
		Password      string `yaml:"password"`
		EnableLogging bool   `yaml:"logging"`
	} `yaml:"server"`
}

func getUser() string {
	var c Config
	if err := cleanenv.ReadConfig("conf/conf.yaml", &c); err != nil {
		processError(err)
	}
	user := c.Server.User
	return user
}

func getPass() string {
	var c Config
	if err := cleanenv.ReadConfig("conf/conf.yaml", &c); err != nil {
		processError(err)
	}
	pass := c.Server.Password
	return pass
}

// TODO make unified function getConfigValue() to return specific property name
func GetConfigValue() {
	return
}
